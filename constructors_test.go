package josh_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orsinium-labs/josh"
)

func TestOk(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Resp[int] {
		return josh.Ok(13)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 200)
	b := must(io.ReadAll(resp.Body))
	eq(string(b), `{"data":13}`+"\n")
}

func TestCreated(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Resp[int] {
		return josh.Created(13)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 201)
	b := must(io.ReadAll(resp.Body))
	eq(string(b), `{"data":13}`+"\n")
}

func TestAccepted(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Resp[int] {
		return josh.Accepted(13)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 202)
	b := must(io.ReadAll(resp.Body))
	eq(string(b), `{"data":13}`+"\n")
}

func TestNoContent(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		return josh.NoContent[josh.Z]()
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 204)
	eq(resp.ContentLength, 0)
}

func TestNotModified(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		return josh.NotModified[josh.Z]()
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 304)
	eq(resp.ContentLength, 0)
}

func TestBadRequest(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		err := josh.Error{Detail: "oh no"}
		return josh.BadRequest[josh.Z](err)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 400)
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"errors":[{"detail":"oh no"}]}`+"\n")
}

func TestUnauthorized(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		err := josh.Error{Detail: "oh no"}
		return josh.Unauthorized[josh.Z](err)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 401)
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"errors":[{"detail":"oh no"}]}`+"\n")
}

func TestForbidden(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		err := josh.Error{Detail: "oh no"}
		return josh.Forbidden[josh.Z](err)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 403)
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"errors":[{"detail":"oh no"}]}`+"\n")
}

func TestNotFound(t *testing.T) {
	h := josh.Wrap(func(r *http.Request) josh.Void {
		err := josh.Error{Detail: "oh no"}
		return josh.NotFound[josh.Z](err)
	})
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	h(w, req)
	resp := w.Result()
	eq(resp.StatusCode, 404)
	body := must(io.ReadAll(resp.Body))
	eq(string(body), `{"errors":[{"detail":"oh no"}]}`+"\n")
}
