package josh_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/orsinium-labs/josh"
)

func TestGetID(t *testing.T) {
	h := josh.Wrap(func(r josh.Req) josh.Resp {
		id, errResp := josh.GetID[int](r, "id")
		if errResp != nil {
			return josh.BadRequest(*errResp)
		}
		return josh.Ok(id)
	})

	{
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		h(w, req)
		resp := w.Result()
		eq(resp.StatusCode, 400)
		body := must(io.ReadAll(resp.Body))
		eq(string(body), `{"errors":[{"detail":"path parameter id is required"}]}`+"\n")
	}

	{
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		req.SetPathValue("id", "13")
		w := httptest.NewRecorder()
		h(w, req)
		resp := w.Result()
		eq(resp.StatusCode, 200)
		body := must(io.ReadAll(resp.Body))
		eq(string(body), `{"data":13}`+"\n")
	}

	{
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		req.SetPathValue("id", "hi")
		w := httptest.NewRecorder()
		h(w, req)
		resp := w.Result()
		eq(resp.StatusCode, 400)
		body := must(io.ReadAll(resp.Body))
		eq(string(body), `{"errors":[{"detail":"invalid path parameter id"}]}`+"\n")
	}

}
