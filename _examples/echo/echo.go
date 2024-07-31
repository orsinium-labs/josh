// The example shows how to make a simple echo server.
//
// Start the server:
//
//	go run ./_examples/echo
//
// Send a request:
//
//	curl http://localhost:8080 --data-raw '{"data":"hello"}'
//
// It should return the following response:
//
//	{"data":"HELLO"}
//
// Try sending an invalid request:
//
//	curl http://localhost:8080 --data-raw '{"data":13}'
//
// Response:
//
//	{"errors":[{"title":"Cannot parse JSON request","detail":"json: cannot unmarshal number into Go struct field .data of type string"}]}
package main

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/josh"
)

func handler(r josh.Req) josh.Resp {
	msg, err := josh.Read[string](r)
	if err != nil {
		return josh.BadRequest(josh.Error{
			Title:  "Cannot parse JSON request",
			Detail: err.Error(),
		})
	}
	msg = strings.ToUpper(msg)
	return josh.Ok(msg)
}

func main() {
	s := josh.NewServer(":8080")
	r := josh.Router{
		"/": {
			GET: josh.Wrap(handler),
		},
	}
	r.Register(nil)
	fmt.Println("listening on http://localhost:8080")
	_ = s.ListenAndServe()
}
