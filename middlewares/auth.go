package middlewares

import (
	"errors"
	"strings"

	"github.com/orsinium-labs/josh"
)

type AuthValidator[U any] func(string) (U, error)

// Authenticate the request using the provided validator before proceeding to the handler.
//
// The validators input is the "Bearer" token provided in the "Authorization" request header.
//
// If the "Authorization" header is not provided, we also check the "Sec-WebSocket-Protocol"
// as a necessary workaround for authenticating WebSocket requests from browsers
// because the browser WebSocket API doesn't allow setting custom headers:
//
// https://github.com/whatwg/websockets/issues/16
//
// The very first protocol after the "Authorization" protocol is treated as
// the authorization token.
//
// If the validator returns an error, that error is immediately returned
// as an "Unathorized" response. Otherwise, the returned value (typically, the user
// or their ID) is added into the request context using [josh.WithSingleton].
func Auth[U any](v AuthValidator[U], h josh.Handler) josh.Handler {
	return func(r josh.Req) josh.Resp {
		user, err := validateRequest(v, r)
		if err != nil {
			return josh.Unauthorized(
				josh.Error{Detail: err.Error()},
			)
		}
		r = josh.Must(josh.WithSingleton(r, user))
		return h(r)
	}
}

func validateRequest[U any](validator AuthValidator[U], r josh.Req) (U, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		// If Authorization header is not provided,
		// try authenticating it as a WebSocket request.
		return validateWSRequest(validator, r)
	}
	token, hasPrefix := strings.CutPrefix(header, "Bearer ")
	if !hasPrefix {
		var def U
		return def, errors.New("Unsupported Authorization type")
	}
	token = strings.TrimSpace(token)
	if token == "" {
		var def U
		return def, errors.New("Authorization token is empty")
	}
	return validator(token)
}

func validateWSRequest[U any](validator AuthValidator[U], r josh.Req) (U, error) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Protocol
	headers := r.Header.Values("Sec-WebSocket-Protocol")
	var def U
	if len(headers) == 0 {
		return def, errors.New("Authorization header not found")
	}

	foundAuth := false
	for _, tokens := range headers {
		for _, token := range strings.Split(tokens, ",") {
			token = strings.TrimSpace(token)
			if token == "Authorization" {
				foundAuth = true
			} else if foundAuth {
				return validator(token)
			}
		}
	}

	return def, errors.New("Authorization header not found")
}
