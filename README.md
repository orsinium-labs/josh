# josh

[ [📚 docs](https://pkg.go.dev/github.com/orsinium-labs/josh) ] [ [🐙 github](https://github.com/orsinium-labs/josh) ]

Go library making it safe and easy to write [JSON:API](https://jsonapi.org/) services. You could call it a we framework except (unlike other web frameworks) josh is very minimalistic. It works together with [net/http](https://pkg.go.dev/net/http) and doesn't reimplement anything that Go standard library already does well.

* Takes care of JSON response serialization and request deserialization.
* Follows JSON:API spec.
* Embraces semantic status codes.
* Small an easy to learn.
* Zero dependency.
* Modular and works great with stdlib or any other web framework.
* Uses the latest Go features: request router, structured logging, generics.
* Statically ensures that you don't make common mistakes:
  * Don't forget the status code.
  * Don't set body for statuses that don't support body.
  * Don't try sending more headers when headers are already sent.

## 🤷 Why

* **Why yet another web framework?** The standard library [net/http](https://pkg.go.dev/net/http) is very generic. You can make writing APIs much faster and safer if you assume that most of endpoints consume and return JSON. Third-party web frameworks (like [gin](https://github.com/gin-gonic/gin)) all do that but they all are unnecessary bloated. There is no need to reimplement multiplexer, logger, or request struct if what Go has out of the box is already very good. Also, you can make a better API with generics which are relatively new to the language and not adopted by any of the old frameworks.
* **Why REST?** [GraphQL](https://graphql.org/) is great but you get much better performance and safety if the backend is in full control of what data is fetched from the database and how. [gRPC](https://grpc.io/) is also amazing but doesn't work with frontend. You could use [grpc-web](https://github.com/grpc/grpc-web) but that's just a proxy that converts a web-compatible API (kinda REST) into gRPC and back, whcih adds unncessary complexity to the system.
* **Why JSON:API?** JSON:API is the most well-adopted and well-documented spec. You could reinvent your own specification (which most companies do) but following a specification written but other smart people saves you from a lot of mistakes and lets you to spend your time on focusing on other important things.
* **Why JSON?** That's the only serialization format supported by frontend out of the box. Anything else requires third-party NPM dependencies, which bloats the frontend app bundle size and introduces security risks. Also, JSON is the most widely supported format with a lot of tooling build around it.

## 📦 Installation

```bash
github.com/orsinium-labs/josh
```

## 🔧 Usage

Here is a simple handler:

```go
func handler(r *http.Request) josh.Resp[string] {
  // Read JSON request body
  msg, err := josh.Read[string](r)
  if err != nil {
    // Return an error for invalid request
    return josh.BadRequest[string](josh.Error{
      Title:  "Cannot parse JSON request",
      Detail: err.Error(),
    })
  }
  // Return the uppercase message
  msg = strings.ToUpper(msg)
  return josh.Ok(msg)
}
```

And here is a server using it:

```go
func main() {
  s := josh.NewServer(":8080")
  r := josh.Router{
    "/": {
      GET: josh.Wrap(handler),
    },
  }
  r.Register(nil)
  _ = s.ListenAndServe()
}
```

Or if you don't want to use the router and the custom server:

```go
func main() {
  http.ListenAndServe(":8080", josh.Wrap(handler))
}
```
