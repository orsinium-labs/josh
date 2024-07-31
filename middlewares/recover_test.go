package middlewares_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/orsinium-labs/josh"
	"github.com/orsinium-labs/josh/middlewares"
)

func eq[T comparable](a, b T) {
	if a != b {
		panic(fmt.Sprintf("%v != %v", a, b))
	}
}

func TestRecover_Ok(t *testing.T) {
	hf := func(r josh.Req) josh.Resp {
		return josh.Ok("hi")
	}
	h := josh.Wrap(middlewares.Recover(hf))
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	body := josh.Must(io.ReadAll(resp.Body))
	eq(string(body), `{"data":"hi"}`+"\n")
}

func TestRecover_Panic(t *testing.T) {
	hf := func(r josh.Req) josh.Resp {
		panic("oh no")
	}
	h := josh.Wrap(middlewares.Recover(hf))
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 500)
}
