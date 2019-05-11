// Endpoints act as proxies for the business functionality. In order to do that one needs to shrink and abstract
// away the parameters. Enter request_response.go...
package endpoint

type (
	HelloRequest struct {
		Name string `json:"name"`
	}
	HelloResponse struct {
		Greeting string `json:"greeting"`
		Err      error  `json:"-"`
	}
)

// In order to determine whether an endpoint failed, we need to implement the Failer interface.
// That way we can use middlewares to act on errors
func (r HelloResponse) Failed() error { return r.Err }

