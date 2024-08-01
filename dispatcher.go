package josh

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
)

type Dispatcher struct {
	handlers map[string]func(context.Context, json.RawMessage) Resp
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		handlers: make(map[string]func(context.Context, json.RawMessage) Resp),
	}
}

func Register[R any](d *Dispatcher, t string, h func(context.Context, R) Resp) {
	if d.handlers == nil {
		panic("josh.Dispatcher must be constructed using josh.NewDispatcher")
	}
	_, exists := d.handlers[t]
	if exists {
		panic("the Dispatcher already contains handler for the given type")
	}
	d.handlers[t] = func(ctx context.Context, raw json.RawMessage) Resp {
		decoder := json.NewDecoder(bytes.NewBuffer(raw))
		decoder.DisallowUnknownFields()
		var req R
		err := decoder.Decode(&req)
		if err != nil {
			// TODO: better error message
			return BadRequest(Error{
				Title:  "Invalid JSON request",
				Detail: err.Error(),
			})
		}
		return h(ctx, req)
	}
}

func (d *Dispatcher) Read(ctx context.Context, r io.Reader) Resp {
	envelope := struct {
		Data *Data[json.RawMessage] `json:"data"`
	}{}
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&envelope)
	if err != nil {
		// TODO: better error message
		return BadRequest(Error{
			Title:  "Invalid JSON request",
			Detail: err.Error(),
		})

	}
	if envelope.Data == nil {
		return BadRequest(Error{
			Title: "JSON request misses the data field",
		})
	}
	if envelope.Data.ID != "" {
		return BadRequest(Error{
			Title: "request cannot contain id",
		})
	}
	h, found := d.handlers[envelope.Data.Type]
	if !found {
		return BadRequest(Error{
			Title: "Unsupported request type",
		})
	}
	ctx = patchLogger(ctx, envelope.Data.Type)
	return h(ctx, envelope.Data.Attributes)
}

// If the context has a logger, add the request-type into the log extras.
func patchLogger(ctx context.Context, t string) context.Context {
	raw := ctx.Value(ctxKey[*slog.Logger]{})
	if raw == nil {
		return ctx
	}
	logger := raw.(*slog.Logger)
	logger = logger.With("request-type", t)
	return context.WithValue(ctx, ctxKey[*slog.Logger]{}, logger)
}
