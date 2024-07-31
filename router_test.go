package josh_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orsinium-labs/josh"
)

func noop(*http.Request) josh.Resp {
	return josh.NoContent()
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func eq[T comparable](a, b T) {
	if a != b {
		panic(fmt.Sprintf("%v != %v", a, b))
	}
}

func TestRouter(t *testing.T) {
	r := josh.Router{
		"/post": {
			GET: josh.Wrap(noop),
		},
	}
	mux := http.NewServeMux()
	r.Register(mux)
	s := httptest.NewServer(mux)
	defer s.Close()

	eq(must(http.Get(s.URL)).StatusCode, 404)
	// right path, right method
	eq(must(http.Get(s.URL+"/post")).StatusCode, 204)
	// wrong method
	eq(must(http.Post(s.URL+"/post", "", nil)).StatusCode, 405)
	// shorter path
	eq(must(http.Get(s.URL+"/pos")).StatusCode, 404)
	// longer path
	eq(must(http.Get(s.URL+"/posts")).StatusCode, 404)
	// subpaths
	eq(must(http.Get(s.URL+"/post/1")).StatusCode, 404)
}
