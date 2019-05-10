package grpc

import (
	"context"

	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
	"errors"
)

func EncodeError(err error) error {
	return err
}

func EncodeHelloRequest(ctx context.Context, request interface{}) (pbRequest interface{}, err error) {
	if pbRequest == nil {
		return nil, errors.New("nil HelloRequest")
	}
	req := request.(endpoint.HelloRequest)
	pbRequest = &example.HelloRequest{
		Name: req.Name,
	}

	return pbRequest, nil
}

func DecodeHelloRequest(ctx context.Context, pbRequest interface{}) (request interface{}, err error) {
	if pbRequest == nil {
		return nil, errors.New("nil HelloRequest")
	}
	req := pbRequest.(*example.HelloRequest)
	request = endpoint.HelloRequest{
		Name: req.Name,
	}
	return request, nil
}

func EncodeHelloResponse(ctx context.Context, response interface{}) (pbResponse interface{}, err error) {
	if response == nil {
		return nil, errors.New("nil HelloResponse")
	}
	resp := response.(endpoint.HelloResponse)
	pb := &example.HelloResponse{
		Greeting: &example.Greeting{
			Text: resp.Greeting,
		},
	}

	return pb, nil
}

func DecodeHelloResponse(ctx context.Context, pbResponse interface{}) (response interface{}, err error) {
	if response == nil {
		return nil, errors.New("nil HelloResponse")
	}
	resp := pbResponse.(*example.HelloResponse)
	response = endpoint.HelloResponse{
		Greeting: resp.Greeting.Text,
	}

	return response, nil
}

