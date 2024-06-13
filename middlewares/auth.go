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
// If the validator returns an error, that erro is immediately returned
// as an "Unathorized" response. Otherwise, the returned value (typically, the user
// or their ID) is added into the request context using [josh.WithSingleton].
func Auth[U, R any](v AuthValidator[U], h josh.Handler[R]) josh.Handler[R] {
	return func(r josh.Req) josh.Resp[R] {
		user, err := validateRequest(v, r)
		if err != nil {
			return josh.Unauthorized[R](
				josh.Error{Detail: err.Error()},
			)
		}
		r = josh.Must(josh.WithSingleton(r, user))
		return h(r)
	}
}

func validateRequest[U any](validator AuthValidator[U], r josh.Req) (U, error) {
	header := r.Header.Get("Authorization")
	var def U
	if header == "" {
		return def, errors.New("Authorization header not found")
	}
	token, hasPrefix := strings.CutPrefix(header, "Bearer ")
	if !hasPrefix {
		return def, errors.New("Unsupported Authorization type")
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return def, errors.New("Authorization token is empty")
	}
	return validator(token)
}
