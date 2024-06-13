package middlewares

import (
	"github.com/orsinium-labs/josh"
)

// Middleware that catches panics, recovers, logs them, and returns 500.
func Recover[T any](h josh.Handler[T]) josh.Handler[T] {
	return func(r josh.Req) (resp josh.Resp[T]) {
		defer func() {
			err := recover()
			if err != nil {
				errResp := josh.Error{Title: "Internal server error"}
				resp = josh.InternalServerError[T](errResp)
			}
		}()
		return h(r)
	}
}
