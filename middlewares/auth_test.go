package middlewares_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orsinium-labs/josh"
	"github.com/orsinium-labs/josh/middlewares"
)

func TestAuth(t *testing.T) {
	var req josh.Req
	url := "http://example.com/foo"

	req = httptest.NewRequest("GET", url, nil)
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "secret")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer ohno")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer secret")
	checkAuth(t, req, 200)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "secret")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "secret, Authorization")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "Authorization, ohno")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "Authorization, secret")
	checkAuth(t, req, 200)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "Authorization")
	req.Header.Add("Sec-WebSocket-Protocol", "secret")
	checkAuth(t, req, 200)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "secret")
	req.Header.Add("Sec-WebSocket-Protocol", "Authorization")
	checkAuth(t, req, 401)

	req = httptest.NewRequest("GET", url, nil)
	req.Header.Add("Sec-WebSocket-Protocol", "Authorization")
	req.Header.Add("Sec-WebSocket-Protocol", "ohno")
	req.Header.Add("Sec-WebSocket-Protocol", "secret")
	checkAuth(t, req, 401)
}

func checkAuth(t *testing.T, req *http.Request, code int) {
	t.Helper()
	type User string
	h := func(r josh.Req) josh.Resp {
		u, err := josh.GetSingleton[User](r)
		if err != nil {
			t.Fatal(err)
		}
		if u != "Aragorn" {
			t.Fatalf("got %s, expected Aragorn", u)
		}
		return josh.Ok("all is good")
	}
	v := func(token string) (User, error) {
		if token == "secret" {
			return "Aragorn", nil
		}
		return "", fmt.Errorf("bad token: %s", token)
	}
	h = middlewares.Auth(v, h)
	hh := josh.Wrap(h)

	w := httptest.NewRecorder()
	hh(w, req)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	t.Log(string(body))
	if resp.StatusCode != code {
		t.Fatalf("got %d, expected %d", resp.StatusCode, code)
	}
}
