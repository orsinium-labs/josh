package josh

import "github.com/orsinium-labs/josh/statuses"

// Do not send any response.
//
// Return it from a handler when the handler handles sanding the response itself.
// For example, when it establishes a WebSocket connection.
func NoResponse() Resp {
	return Resp{}
}

// Respond with 200 status code.
//
// https://jsonapi.org/format/#fetching-resources-responses-200
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
func Ok(v any) Resp {
	return Resp{
		Status: statuses.OK,
		Data:   &v,
	}
}

// Respond with 201 status code.
func Created(v any) Resp {
	return Resp{
		Status: statuses.Created,
		Data:   &v,
	}
}

// Respond with 202 status code.
func Accepted(v any) Resp {
	return Resp{
		Status: statuses.Accepted,
		Data:   &v,
	}
}

// Respond with 204 status code.
//
// Indicates that there is no content to send for this request.
//
// The response does not have a body.
func NoContent() Resp {
	return Resp{
		Status: statuses.NoContent,
	}
}

// Respond with 304 status code.
//
// Tells the client that the response has not been modified,
// so the client can continue to use the same cached version of the response.
//
// The response does not have a body.
func NotModified() Resp {
	return Resp{
		Status: statuses.NotModified,
	}
}

// Respond with 400 status code.
func BadRequest(err Error) Resp {
	return Resp{
		Status: statuses.BadRequest,
		Errors: []Error{err},
	}
}

// Respond with 401 status code.
func Unauthorized(err Error) Resp {
	return Resp{
		Status: statuses.Unauthorized,
		Errors: []Error{err},
	}
}

// Respond with 403 status code.
func Forbidden(err Error) Resp {
	return Resp{
		Status: statuses.Forbidden,
		Errors: []Error{err},
	}
}

// Respond with 404 status code.
//
// https://jsonapi.org/format/#fetching-resources-responses-404
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
func NotFound(err Error) Resp {
	return Resp{
		Status: statuses.NotFound,
		Errors: []Error{err},
	}
}

// Respond with 500 status code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
func InternalServerError(err Error) Resp {
	return Resp{
		Status: statuses.InternalServerError,
		Errors: []Error{err},
	}
}
