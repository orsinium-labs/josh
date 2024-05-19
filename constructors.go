package josh

import "github.com/orsinium-labs/josh/statuses"

func Ok[T any](v T) Resp[T] {
	return Resp[T]{
		Status:  statuses.OK,
		Content: v,
	}
}
