package josh

import (
	"bytes"
	"encoding/json"
	"io"
)

type Dispatcher struct {
	handlers map[string]func(json.RawMessage) Resp
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		handlers: make(map[string]func(json.RawMessage) Resp),
	}
}

func Register[R any](d *Dispatcher, t string, h func(R) Resp) {
	if d.handlers == nil {
		panic("josh.Dispatcher must be constructed using josh.NewDispatcher")
	}
	_, exists := d.handlers[t]
	if exists {
		panic("the Dispatcher already contains handler for the given type")
	}
	d.handlers[t] = func(raw json.RawMessage) Resp {
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
		return h(req)
	}
}

func (d *Dispatcher) HandleRequest(r Req) Resp {
	return d.HandleReader(r.Body)
}

func (d *Dispatcher) HandleReader(r io.Reader) Resp {
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
	h, found := d.handlers[envelope.Data.Type]
	if !found {
		return BadRequest(Error{
			Title: "Unsupported request type",
		})
	}
	return h(envelope.Data.Attributes)
}
