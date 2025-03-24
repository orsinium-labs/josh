package sse

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/orsinium-labs/josh"
)

type Stream struct {
	req     josh.Req
	flusher http.Flusher
	writer  http.ResponseWriter
	started bool
}

// Create a new SSE helper from a request.
func New(req josh.Req) Stream {
	writer := josh.Must(josh.GetSingleton[http.ResponseWriter](req))
	flusher, ok := writer.(http.Flusher)
	if !ok {
		panic("streaming is disabled on the server")
	}
	return Stream{
		req:     req,
		flusher: flusher,
		writer:  writer,
	}
}

func (s *Stream) Start() {
	s.writer.Header().Set("Content-Type", "text/event-stream")
	s.writer.Header().Set("Cache-Control", "no-cache")
	s.writer.Header().Set("Connection", "keep-alive")
	s.writer.WriteHeader(http.StatusOK)
	s.flusher.Flush()
	s.started = true
}

// Send a success message to the client. Typically, a [josh.Data] instance.
func (s *Stream) SendOk(data any) error {
	_, isMsg := data.(Message)
	if isMsg {
		return errors.New("use Send to send Message")
	}
	msg := Message{Data: josh.Ok(data)}
	return s.Send(msg)
}

// Send an error message to the client.
func (s *Stream) SendError(data josh.Error) error {
	msg := Message{Data: josh.BadRequest(data)}
	return s.Send(msg)
}

// Send a message with additional attributes.
//
// Use [SendOk] or [SendError] if the only field you use is Data.
func (s *Stream) Send(msg Message) error {
	if !s.started {
		return errors.New("you must Start SSE connection before you can Send")
	}
	var resErr error
	w := func(data []byte) {
		_, writeErr := s.writer.Write(data)
		if resErr == nil && writeErr != nil {
			resErr = writeErr
		}
	}

	if msg.Event != "" {
		w([]byte("event: "))
		w([]byte(msg.Event))
		w([]byte{'\n'})
	}
	if msg.Data.Status != 0 {
		w([]byte("data: "))
		raw, err := json.Marshal(msg.Data)
		if err != nil {
			return err
		}
		w(raw)
		w([]byte{'\n'})
	}
	if msg.ID != "" {
		w([]byte("id: "))
		w([]byte(msg.ID))
		w([]byte{'\n'})
	}
	if msg.Retry != 0 {
		w([]byte("retry: "))
		reconn := strconv.FormatInt(msg.Retry.Milliseconds(), 10)
		w([]byte(reconn))
		w([]byte{'\n'})
	}
	if msg.Comment != "" {
		w([]byte{':', ' '})
		w([]byte(msg.Comment))
		w([]byte{'\n'})
	}

	w([]byte{'\n'})
	s.flusher.Flush()
	return resErr
}

type Message struct {
	// A string identifying the type of event described.
	//
	// If this is specified, an event will be dispatched on the browser
	// to the listener for the specified event name; the website source code
	// should use addEventListener() to listen for named events.
	// The onmessage handler is called if no event name is specified for a message.
	Event string

	// The data field for the message. Will be JSON-encoded.
	Data josh.Resp

	// The event ID to set the EventSource object's last event ID value.
	ID string

	// The reconnection time. If the connection to the server is lost,
	// the browser will wait for the specified time before attempting to reconnect.
	Retry time.Duration

	// The comment line can be used to prevent connections from timing out;
	// a server can send a comment periodically to keep the connection alive.
	Comment string
}
