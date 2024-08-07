package middlewares

import (
	"github.com/orsinium-labs/josh"
)

// With is a middleware that adds the given value into request context using [josh.WithSingleton].
func With[D any](data D, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, data))
		return h(r)
	}
}

// With2 is like [With] but sets 2 values in one go.
func With2[A, B any](a A, b B, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, a))
		r = josh.Must(josh.WithSingleton(r, b))
		return h(r)
	}
}

// With3 is like [With] but sets 3 values in one go.
func With3[A, B, C any](a A, b B, c C, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, a))
		r = josh.Must(josh.WithSingleton(r, b))
		r = josh.Must(josh.WithSingleton(r, c))
		return h(r)
	}
}

// With4 is like [With] but sets 4 values in one go.
func With4[A, B, C, D any](a A, b B, c C, d D, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, a))
		r = josh.Must(josh.WithSingleton(r, b))
		r = josh.Must(josh.WithSingleton(r, c))
		r = josh.Must(josh.WithSingleton(r, d))
		return h(r)
	}
}

// With5 is like [With] but sets 5 values in one go.
func With5[A, B, C, D, E any](a A, b B, c C, d D, e E, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, a))
		r = josh.Must(josh.WithSingleton(r, b))
		r = josh.Must(josh.WithSingleton(r, c))
		r = josh.Must(josh.WithSingleton(r, d))
		r = josh.Must(josh.WithSingleton(r, e))
		return h(r)
	}
}

// With5 is like [With] but sets 6 values in one go.
func With6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F, h josh.Handler) josh.Handler {
	return func(r josh.Req) (resp josh.Resp) {
		r = josh.Must(josh.WithSingleton(r, a))
		r = josh.Must(josh.WithSingleton(r, b))
		r = josh.Must(josh.WithSingleton(r, c))
		r = josh.Must(josh.WithSingleton(r, d))
		r = josh.Must(josh.WithSingleton(r, e))
		r = josh.Must(josh.WithSingleton(r, f))
		return h(r)
	}
}
