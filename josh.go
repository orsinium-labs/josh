package josh

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orsinium-labs/josh/statuses"
)

type contextKey string
type literal string

const headersKey contextKey = "headers"

// Handler function type. Accepts a request, returns a response.
type Handler[T any] func(*http.Request) Resp[T]

// Wrap a [Handler] function to make it compatible with stdlib [http.HandlerFunc].
func Wrap[T any](h Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), headersKey, w.Header())
		r = r.WithContext(ctx)
		resp := h(r)
		resp.Write(w)
	}
}

// Set a response header.
func SetHeader(r *http.Request, key literal, value string) {
	headers := r.Context().Value(headersKey).(http.Header)
	headers.Set(string(key), value)
}

// Resp is a response type.
//
// The generic type T is the type of the data response.
type Resp[T any] struct {
	// The response status code.
	Status  statuses.Status
	Content T
	Errors  []Error
}

// Void is a type alias for Resp for when the endpoint does not return any data ever.
//
// The endpoint still can return [NoContent], [NotModified], or an error.
type Void = Resp[struct{}]

// Write response into the connection.
func (r Resp[T]) Write(w http.ResponseWriter) {
	// Write content type and status code.
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/vnd.api+json")
	}
	// If status code allows for body, write the JSON response.
	if !bodyAllowedForStatus(r.Status) {
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
	}{r.Content}
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
