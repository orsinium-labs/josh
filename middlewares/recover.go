package middlewares

import (
	"log/slog"

	"github.com/orsinium-labs/josh"
)

// Middleware that catches panics, recovers, logs them, and returns 500.
func Recover[T any](h josh.Handler[T]) josh.Handler[T] {
	return func(r josh.Req) (resp josh.Resp[T]) {
		defer func() {
			err := recover()
			if err != nil {
				logger, _ := josh.GetSingleton[*slog.Logger](r)
				if logger != nil {
					logger.ErrorContext(
						r.Context(),
						"panic when handling request",
						"error", err,
						"path", r.URL.Path,
						"method", r.Method,
					)
				}
				errResp := josh.Error{Title: "Internal server error"}
				resp = josh.InternalServerError[T](errResp)
			}
		}()
		return h(r)
	}
}
