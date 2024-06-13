package josh

import "github.com/orsinium-labs/josh/statuses"

// Respond with 200 status code.
//
// https://jsonapi.org/format/#fetching-resources-responses-200
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
func Ok[T any](v T) Resp[T] {
	return Resp[T]{
		Status: statuses.OK,
		Data:   &v,
	}
}

// Respond with 201 status code.
func Created[T any](v T) Resp[T] {
	return Resp[T]{
		Status: statuses.Created,
		Data:   &v,
	}
}

// Respond with 202 status code.
func Accepted[T any](v T) Resp[T] {
	return Resp[T]{
		Status: statuses.Accepted,
		Data:   &v,
	}
}

// Respond with 204 status code.
//
// Indicates that there is no content to send for this request.
//
// The response does not have a body.
func NoContent[T any]() Resp[T] {
	return Resp[T]{
		Status: statuses.NoContent,
	}
}

// Respond with 304 status code.
//
// Tells the client that the response has not been modified,
// so the client can continue to use the same cached version of the response.
//
// The response does not have a body.
func NotModified[T any]() Resp[T] {
	return Resp[T]{
		Status: statuses.NotModified,
	}
}

// Respond with 400 status code.
func BadRequest[T any](err Error) Resp[T] {
	return Resp[T]{
		Status: statuses.BadRequest,
		Errors: []Error{err},
	}
}

// Respond with 401 status code.
func Unauthorized[T any](err Error) Resp[T] {
	return Resp[T]{
		Status: statuses.Unauthorized,
		Errors: []Error{err},
	}
}

// Respond with 403 status code.
func Forbidden[T any](err Error) Resp[T] {
	return Resp[T]{
		Status: statuses.Forbidden,
		Errors: []Error{err},
	}
}

// Respond with 404 status code.
//
// https://jsonapi.org/format/#fetching-resources-responses-404
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
func NotFound[T any](err Error) Resp[T] {
	return Resp[T]{
		Status: statuses.NotFound,
		Errors: []Error{err},
	}
}

// Respond with 500 status code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
func InternalServerError[T any](err Error) Resp[T] {
	return Resp[T]{
		Status: statuses.InternalServerError,
		Errors: []Error{err},
	}
}
