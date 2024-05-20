package josh_test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/orsinium-labs/josh"
)

func TestNewServer(t *testing.T) {
	s := josh.NewServer("example.com:1337")
	eq(s.Addr, "example.com:1337")
	eq(s.ReadTimeout, 5*time.Second)
}

func TestSetHeader(t *testing.T) {
	h := josh.Wrap(func(r josh.Req) josh.Void {
		josh.SetHeader(r, "X-Test", "hello")
		return josh.NoContent[josh.Z]()
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 204)
	eq(resp.Header.Get("X-Test"), "hello")
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
}

func TestRead(t *testing.T) {
	h := josh.Wrap(func(r josh.Req) josh.Resp[string] {
		msg, err := josh.Read[string](r)
		if err != nil {
			panic(err)
		}
		msg = strings.ToUpper(msg)
		return josh.Ok(msg)
	})
	req := httptest.NewRequest(
		"GET", "http://example.com/foo",
		bytes.NewReader([]byte(`{"data":"hello"}`)),
	)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"data":"HELLO"}`+"\n")
}
