package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
)

// HelloEndpoint constructs the endpoint of the Hello() method
func HelloEndpoint(service service.Example) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(HelloRequest)
		res, err := service.Hello(ctx, req.Name)
		return HelloResponse{res, err}, err
	}
}
