package josh_test

import (
	"bytes"
	"io"
	"net/http"
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
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		josh.SetHeader(r, "X-Test", "hello")
		return josh.NoContent()
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 204)
	eq(resp.Header.Get("X-Test"), "hello")
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
}

// It should be possible to unwrap http.ResponseWriter and write the response directly.
func TestUnwrapResponseWriter(t *testing.T) {
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		w := josh.Must(josh.GetSingleton[http.ResponseWriter](r))
		w.WriteHeader(204)
		return josh.Resp{}
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 204)
	eq(resp.Header.Get("Content-Type"), "")
}

func TestUnwrap(t *testing.T) {
	hh := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(203)
	}
	h := josh.Wrap(josh.Unwrap(hh))
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 203)
	eq(resp.Header.Get("Content-Type"), "")
}

func TestRead(t *testing.T) {
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		body, err := josh.Read[string]("foo", r.Body)
		if err != nil {
			panic(err)
		}
		msg := body.Attributes
		msg = strings.ToUpper(msg)
		return josh.Ok(msg)
	})
	req := httptest.NewRequest(
		"GET", "http://example.com/foo",
		bytes.NewReader([]byte(`{"data":{"type": "foo", "attributes": "hello"}}`)),
	)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"data":"HELLO"}`+"\n")
}

func TestSingleton(t *testing.T) {
	type User struct{ name string }
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		var err error
		_, err = josh.GetSingleton[User](r)
		if err == nil {
			t.FailNow()
		}
		r = josh.Must(josh.WithSingleton(r, User{"aragorn"}))
		user := josh.Must(josh.GetSingleton[User](r))
		eq(user.name, "aragorn")
		return josh.Ok("ok")
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"data":"ok"}`+"\n")
}

func TestGetOrSetSingleton(t *testing.T) {
	type User struct{ name string }
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		r, user := josh.GetOrSetSingleton(r, func() User {
			return User{"aragorn"}
		})
		eq(user.name, "aragorn")
		_, user = josh.GetOrSetSingleton(r, func() User {
			return User{"gandalf"}
		})
		eq(user.name, "aragorn")
		return josh.Ok("ok")
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	eq(resp.Header.Get("Content-Type"), "application/vnd.api+json")
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"data":"ok"}`+"\n")
}
