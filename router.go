package josh

import "net/http"

// Router is a mapping of URL paths to endpoints.
type Router map[string]Endpoint

// Register all endpoints from the router in the given stdlib multiplexer.
//
// Pass nil argument if you want to use the default global multiplexer.
func (r Router) Register(mux *http.ServeMux) {
	if mux == nil {
		mux = http.DefaultServeMux
	}
	for path, endpoint := range r {
		if endpoint.GET != nil {
			mux.HandleFunc("GET "+path, endpoint.GET)
		}
		if endpoint.HEAD != nil {
			mux.HandleFunc("HEAD "+path, endpoint.HEAD)
		}
		if endpoint.POST != nil {
			mux.HandleFunc("POST "+path, endpoint.POST)
		}
		if endpoint.PUT != nil {
			mux.HandleFunc("PUT "+path, endpoint.PUT)
		}
		if endpoint.DELETE != nil {
			mux.HandleFunc("DELETE "+path, endpoint.DELETE)
		}
		if endpoint.CONNECT != nil {
			mux.HandleFunc("CONNECT "+path, endpoint.CONNECT)
		}
		if endpoint.OPTIONS != nil {
			mux.HandleFunc("OPTIONS "+path, endpoint.OPTIONS)
		}
		if endpoint.TRACE != nil {
			mux.HandleFunc("TRACE "+path, endpoint.TRACE)
		}
		if endpoint.PATCH != nil {
			mux.HandleFunc("PATCH "+path, endpoint.PATCH)
		}
	}
}

// Endpoint is a collection of handlers for the same path but different HTTP methods.
type Endpoint struct {
	// GET method requests a representation of the specified resource.
	// Requests using GET should only retrieve data.
	//
	// https://jsonapi.org/format/#fetching-resources
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET
	GET http.HandlerFunc

	// HEAD method asks for a response identical to a GET request,
	// but without the response body.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD
	HEAD http.HandlerFunc

	// POST method submits an entity to the specified resource,
	// often causing a change in state or side effects on the server.
	//
	// https://jsonapi.org/format/#crud-creating
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST
	POST http.HandlerFunc

	// PUT method replaces all current representations of the target resource
	// with the request payload.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT
	PUT http.HandlerFunc

	// DELETE method deletes the specified resource.
	//
	// https://jsonapi.org/format/#crud-deleting
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE
	DELETE http.HandlerFunc

	// CONNECT method establishes a tunnel to the server identified by the target resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT
	CONNECT http.HandlerFunc

	// OPTIONS method describes the communication options for the target resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/OPTIONS
	OPTIONS http.HandlerFunc

	// TRACE method performs a message loop-back test along the path to the target resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/TRACE
	TRACE http.HandlerFunc

	// PATCH method applies partial modifications to a resource.
	//
	// https://jsonapi.org/format/#crud-updating
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH
	PATCH http.HandlerFunc
}
