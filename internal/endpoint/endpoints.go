// Here we define all endpoints, including middlewares, which are exposed via the service interface.
package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
	"context"
)

// Encpoints initializes the Set with all endpoints including their middleware
func Endpoints(svc service.ExampleService) Set {

	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = HelloEndpoint(svc)
	}

	return Set{
		HelloEndpoint: helloEndpoint,
	}
}

// HelloEndpoint constructs the endpoint of the Hello() method
func HelloEndpoint(service service.ExampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HelloRequest)
		res, err := service.Hello(ctx, req.Name)
		return HelloResponse{res, err}, err
	}
}
