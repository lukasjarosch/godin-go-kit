package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
)

// Set bundles all of the service's endpoints. Also, Set implements the ExampleService interface.
// That way a Set can also be used just like the service which get's handy when you need access to the service while
// keeping all applied middlewares. This is particularly useful if you want to trigger business logic from a consumer,
// skipping the transport layer.
type Set struct {
	HelloEndpoint endpoint.Endpoint
}

// Endpoints initializes the Set with all endpoints including their middleware
func Endpoints(svc service.Example) Set {

	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = HelloEndpoint(svc)
	}

	return Set{
		HelloEndpoint: helloEndpoint,
	}
}

func (s Set) Hello(ctx context.Context, name string) (greeting string, err error) {
	resp, err := s.HelloEndpoint(ctx, HelloRequest{
		Name: name,
	})
	if err != nil {
		return "", err
	}
	response := resp.(HelloResponse)
	return response.Greeting, response.Err
}
