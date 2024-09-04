package middlewares

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/josh"
	"github.com/orsinium-labs/josh/statuses"
)

// Require the request to have a specific Content-Type.
//
// If "ct" is an empty string, the "application/vnd.api+json" will be required.
func ContentType(ct string, h josh.Handler) josh.Handler {
	if ct == "" {
		ct = "application/vnd.api+json"
	}
	return func(r josh.Req) josh.Resp {
		act := r.Header.Get("Content-Type")
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(act, ";")[0]))
		if mediaType != ct {
			return josh.Resp{
				Status: statuses.UnsupportedMediaType,
				Errors: []josh.Error{{
					Title:  "Unsupported Content-Type",
					Detail: fmt.Sprintf("The request must have %s Content-Type", ct),
				}},
			}
		}
		return h(r)
	}
}
