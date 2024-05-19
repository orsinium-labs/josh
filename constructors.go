package josh

import "github.com/orsinium-labs/josh/statuses"

// Respond with 200 status code.
func Ok[T any](v T) Resp[T] {
	return Resp[T]{
		Status:  statuses.OK,
		Content: v,
	}
}

// Respond with 201 status code.
func Created[T any](v T) Resp[T] {
	return Resp[T]{
		Status:  statuses.Created,
		Content: v,
	}
}

// Respond with 202 status code.
func Accepted[T any](v T) Resp[T] {
	return Resp[T]{
		Status:  statuses.Accepted,
		Content: v,
	}
}

// Respond with 204 status code.
//
// Indicates that there is no content to send for this request.
func NoContent[T any]() Resp[T] {
	return Resp[T]{
		Status: statuses.NoContent,
	}
}

// Respond with 304 status code.
//
// Tells the client that the response has not been modified,
// so the client can continue to use the same cached version of the response.
func NotModified[T any]() Resp[T] {
	return Resp[T]{
		Status: statuses.NotModified,
	}
}
