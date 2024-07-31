package middlewares

import (
	"log/slog"

	"github.com/orsinium-labs/josh"
)

// Middleware that catches panics, recovers, logs them, and returns 500.
func Recover(h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
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
				resp = josh.InternalServerError(errResp)
			}
		}()
		return h(r)
	}
}
