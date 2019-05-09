package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

// Set bundles all of the service's endpoints. Also, Set implements the ExampleService interface.
// That way a Set can also be used just like the service which get's handy when you need access to the service while
// keeping all applied interceptors. This is particularly useful if you want to trigger business logic from a consumer,
// skipping the transport layer.
type Set struct {
	HelloEndpoint endpoint.Endpoint
}

func (Set) Hello(ctx context.Context, name string) (greeting string, err error) {
	return
}
