package josh

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orsinium-labs/josh/statuses"
)

type stringLiteral string

type Handler[T any] func(*http.Request) Resp[T]

func Wrap[T any](h Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)
		resp.Write(w)
	}
}

type Resp[T any] struct {
	Status  statuses.Status
	Content T
	Errors  []Error
}

// Write response into the connection.
func (r Resp[T]) Write(w http.ResponseWriter) {
	// Write content type and status code.
	if w.Header().Get("Content-Type") == "" {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
	}
	w.WriteHeader(int(r.Status))

	// If status code allows for body, write the JSON response.
	if !bodyAllowedForStatus(r.Status) {
		return
	}
	if r.Errors != nil {
		r.writeError(w)
		return
	}
	// TODO: log error
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(r.Content)
}

func (r Resp[T]) writeError(w http.ResponseWriter) {
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
	}{Errors: r.Errors}
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
