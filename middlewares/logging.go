package middlewares

import (
	"log/slog"

	"github.com/orsinium-labs/josh"
)

// Add the logger into the request context and add some request info into all log records.
func WithLogger(logger *slog.Logger, h josh.Handler) josh.Handler {
	return func(req josh.Req) josh.Resp {
		logger = logger.With(
			"method", req.Method,
			"pattern", req.Pattern,
			"path", req.URL.Path,
			"content-length", req.ContentLength,
			"remote-addr", req.RemoteAddr,
		)
		req = josh.Must(josh.WithSingleton(req, logger))
		return h(req)
	}
}
