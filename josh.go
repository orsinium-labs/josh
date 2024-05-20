package josh

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/orsinium-labs/josh/statuses"
)

type contextKey string

const headersKey contextKey = "headers"

// Req is an alias for a pointer to [http.Request].
type Req = *http.Request

// NewServer creates and [http.Server] with safe defaults.
func NewServer(addr string) *http.Server {
	// https://github.com/google/go-safeweb/blob/master/safehttp/server.go#L96
	return &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    10 * 1024,
	}
}

// Handler function type. Accepts a request, returns a response.
type Handler[T any] func(Req) Resp[T]

// Wrap a [Handler] function to make it compatible with stdlib [http.HandlerFunc].
func Wrap[T any](h Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r Req) {
		ctx := context.WithValue(r.Context(), headersKey, w.Header())
		r = r.WithContext(ctx)
		resp := h(r)
		resp.Write(w)
	}
}

// Set a response header.
func SetHeader(r Req, key, value string) {
	headers := r.Context().Value(headersKey).(http.Header)
	headers.Set(key, value)
}

// Read and parse request body as JSON.
func Read[T any](r Req) (T, error) {
	if r.ContentLength == 0 {
		return *new(T), errors.New("request body is empty")
	}
	var v struct {
		Data *T `json:"data"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return *new(T), err
	}
	if v.Data == nil {
		return *new(T), errors.New("data field not found in request body")
	}
	return *v.Data, nil
}

// Resp is a response type.
//
// The generic type T is the type of the data response.
//
// If the handler never returns data, use [Void] instead.
//
// https://jsonapi.org/format/#document-top-level
type Resp[T any] struct {
	// The response status code.
	//
	// If possible, don't create empty responses directly.
	// Instead, use one of the constructors: [NoContent] or [NotModified].
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
	Status statuses.Status `json:"-"`

	// The document's "primary data" returned on success.
	//
	// If possible, don't create sucessful responses directly.
	// Instead, use one of the constructors: [Ok], [Created], or [Accepted].
	//
	// https://jsonapi.org/format/#fetching-resources-responses
	Data T `json:"data,omitempty"`

	// A slice of errors returned on failure.
	//
	// If possible, don't create error responses directly.
	// Instead, use one of the constructors: [BadRequest], [Unauthorized],
	// [Forbidden], or [NotFound].
	//
	// https://jsonapi.org/format/#error-objects
	Errors []Error `json:"errors,omitempty"`
}

// Void is a type alias for [Resp] for when the handler does not return any data ever.
//
// The endpoint still can return [NoContent], [NotModified], or an error.
type Void = Resp[Z]

// Z is a zero type. Used by [Void].
type Z = struct{}

// Write response into the connection.
func (r Resp[T]) Write(w http.ResponseWriter) {
	// Write content type and status code.
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/vnd.api+json")
	}
	// If status code allows for body, write the JSON response.
	if !bodyAllowedForStatus(r.Status) {
		w.Header().Add("Content-Length", "0")
		w.WriteHeader(int(r.Status))
		return
	}
	if r.Errors != nil {
		r.writeErrors(w)
		return
	}
	r.writeData(w)
}

func (r Resp[T]) writeErrors(w http.ResponseWriter) {
	if r.Status == 0 {
		r.Status = statuses.BadRequest
	}
	w.WriteHeader(int(r.Status))
	encoder := json.NewEncoder(w)
	for _, err := range r.Errors {
		if err.Code == "" {
			err.Code = strconv.Itoa(int(r.Status))
		}
		if err.Title == "" {
			err.Title = http.StatusText(int(r.Status))
		}
	}
	// https://jsonapi.org/format/#error-objects
	v := struct {
		Errors []Error `json:"errors"`
	}{r.Errors}
	// TODO: log error
	_ = encoder.Encode(v)
}

func (r Resp[T]) writeData(w http.ResponseWriter) {
	if r.Status == 0 {
		r.Status = statuses.OK
	}
	w.WriteHeader(int(r.Status))
	encoder := json.NewEncoder(w)
	v := struct {
		Data T `json:"data"`
	}{r.Data}
	// TODO: log error
	_ = encoder.Encode(v)
}

func bodyAllowedForStatus(status statuses.Status) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == statuses.NoContent:
		return false
	case status == statuses.NotModified:
		return false
	}
	return true
}
