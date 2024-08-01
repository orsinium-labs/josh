package josh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/orsinium-labs/josh/headers"
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
type Handler func(Req) Resp

// Wrap a [Handler] function to make it compatible with stdlib [http.HandlerFunc].
//
// For going the other way around, see [Unwrap].
func Wrap(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r Req) {
		ctx := context.WithValue(r.Context(), headersKey, w.Header())
		r = r.WithContext(ctx)
		r, _ = WithSingleton(r, w)
		resp := h(r)
		resp.Write(w)
	}
}

// Adapter making [http.HandlerFunc] compatible with josh.
//
// For going the other way around, see [Wrap].
func Unwrap(h http.HandlerFunc) Handler {
	return func(r Req) Resp {
		w := Must(GetSingleton[http.ResponseWriter](r))
		h(w, r)
		return NoResponse()
	}
}

// Set a response header.
func SetHeader(r Req, key headers.Header, value string) {
	headers := r.Context().Value(headersKey).(http.Header)
	headers.Set(string(key), value)
}

// Read and parse request body as JSON.
//
// Accepts the expected value of "type" field and Req.Body.
func Read[T any](t string, r io.Reader) (T, error) {
	// TODO: improve based on this blog post:
	// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	if r == nil {
		return *new(T), errors.New("request body is empty")
	}
	var v struct {
		Data *Data[T] `json:"data"`
	}
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&v)
	if err != nil {
		return *new(T), err
	}
	if v.Data == nil {
		return *new(T), errors.New("data field not found in request body")
	}
	if v.Data.Type != t {
		return *new(T), fmt.Errorf("unexpected request type: expected %s, got %s", t, v.Data.Type)
	}
	if v.Data.ID != "" {
		return *new(T), errors.New("requests cannot contain id")
	}
	return v.Data.Attributes, nil
}

// Resp is a response type.
//
// The generic type T is the type of the data response.
//
// If the handler never returns data, use [Void] instead.
//
// https://jsonapi.org/format/#document-top-level
type Resp struct {
	// The response status code.
	//
	// If possible, don't create empty responses directly.
	// Instead, use one of the constructors: [NoContent] or [NotModified].
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
	Status statuses.Status `json:"-"`

	// The document's "primary data" returned on success.
	//
	// If possible, don't create successful responses directly.
	// Instead, use one of the constructors: [Ok], [Created], or [Accepted].
	//
	// https://jsonapi.org/format/#fetching-resources-responses
	Data any `json:"data,omitempty"`

	// A slice of errors returned on failure.
	//
	// If possible, don't create error responses directly.
	// Instead, use one of the constructors: [BadRequest], [Unauthorized],
	// [Forbidden], or [NotFound].
	//
	// https://jsonapi.org/format/#error-objects
	Errors []Error `json:"errors,omitempty"`

	Included any `json:"included,omitempty"`
	JSONAPI  any `json:"jsonapi,omitempty"`
	Links    any `json:"links,omitempty"`
	Meta     any `json:"meta,omitempty"`
}

// Write response into the connection.
func (r Resp) Write(w http.ResponseWriter) {
	if r.Status == 0 && r.Data == nil && r.Errors == nil {
		return
	}
	// Write content type and status code.
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/vnd.api+json")
	}
	// If status code allows for body, write the JSON response.
	if !bodyAllowedForStatus(r.Status) {
		w.Header().Add("Content-Length", "0")
		if r.Status != 0 {
			w.WriteHeader(int(r.Status))
		}
		return
	}
	if r.Errors != nil {
		r.writeErrors(w)
		return
	}
	r.writeData(w)
}

func (r Resp) writeErrors(w http.ResponseWriter) {
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
			err.Title = r.Status.Text()
		}
	}
	// TODO: log error
	_ = encoder.Encode(r)
}

func (r Resp) writeData(w http.ResponseWriter) {
	if r.Status == 0 {
		r.Status = statuses.OK
	}
	w.WriteHeader(int(r.Status))
	encoder := json.NewEncoder(w)
	// TODO: log error
	_ = encoder.Encode(r)
}

type ctxKey[T any] struct{}

// Wrap a function returning a value and error, panic if the error is not nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Attach the given value to the request's context.
//
// The context can store only one value of the given type.
// If there is already a value of the same type in the context, an error is returned.
// Use [Must] if you're sure that the value is not present.
func WithSingleton[T any](r Req, val T) (Req, error) {
	key := ctxKey[T]{}
	ctx := r.Context()
	if ctx.Value(key) != nil {
		return r, errors.New("context already contains value of the given type")
	}
	ctx = context.WithValue(ctx, key, val)
	return r.WithContext(ctx), nil
}

// Get from the context the value added using [WithSingleton].
//
// If there is no value of the given type in the context, an error is returned.
// Use [Must] if you're sure that the value is present.
func GetSingleton[T any](r Req) (T, error) {
	raw := r.Context().Value(ctxKey[T]{})
	if raw == nil {
		return *new(T), errors.New("no value of the given type in the context")
	}
	return raw.(T), nil
}

// Get singleton from the request context. If not present, build it and put in the context.
//
// Try to [GetSingleton]. If it's not in the request context, make it
// using the passed function and then put it into the request context using [WithSingleton].
//
// It is useful for using the context as cache in scenarios when you want to lazily
// initialize a value that might be accessed from multiple middlewares.
func GetOrSetSingleton[T any](r Req, f func() T) (Req, T) {
	v, err := GetSingleton[T](r)
	if err != nil {
		v = f()
		r, _ = WithSingleton(r, v)
	}
	return r, v
}

// Check if the given status allows a body in the response.
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
