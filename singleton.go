package josh

import (
	"context"
	"errors"
)

// Attach the given value to the request's context.
//
// The context can store only one value of the given type.
// If there is already a value of the same type in the context, an error is returned.
// Use [Must] if you're sure that the value is not present.
func WithSingleton[T any](r Req, val T) (Req, error) {
	key := ctxKey[T]{}
	ctx := r.Context()
	if ctx.Value(key) != nil {
		return r, errors.New("context already contains value of the given type")
	}
	ctx = context.WithValue(ctx, key, val)
	return r.WithContext(ctx), nil
}

// Like [WithSingleton] but operates directly with [context.Context] rather than [Req].
func CWithSingleton[T any](ctx context.Context, val T) (context.Context, error) {
	key := ctxKey[T]{}
	if ctx.Value(key) != nil {
		return ctx, errors.New("context already contains value of the given type")
	}
	ctx = context.WithValue(ctx, key, val)
	return ctx, nil
}

// Get from the context the value added using [WithSingleton].
//
// If there is no value of the given type in the context, an error is returned.
// Use [Must] if you're sure that the value is present.
func GetSingleton[T any](r Req) (T, error) {
	return CGetSingleton[T](r.Context())
}

// Like [GetSingleton] but operates directly with [context.Context] rather than [Req].
func CGetSingleton[T any](ctx context.Context) (T, error) {
	raw := ctx.Value(ctxKey[T]{})
	if raw == nil {
		return *new(T), errors.New("no value of the given type in the context")
	}
	return raw.(T), nil
}

// Get singleton from the request context. If not present, build it and put in the context.
//
// Try to [GetSingleton]. If it's not in the request context, make it
// using the passed function and then put it into the request context using [WithSingleton].
//
// It is useful for using the context as cache in scenarios when you want to lazily
// initialize a value that might be accessed from multiple middlewares.
func GetOrSetSingleton[T any](r Req, f func() T) (Req, T) {
	v, err := GetSingleton[T](r)
	if err != nil {
		v = f()
		r, _ = WithSingleton(r, v)
	}
	return r, v
}
