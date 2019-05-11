package grpc

import (
	"context"

	"errors"

	"github.com/lukasjarosch/godin-go-kit/internal/api"
	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EncodeError(err error) error {
	switch err {
	case example.ErrNotImplemented:
		return status.Error(codes.Unimplemented, err.Error())
	}
	return err
}

func EncodeHelloRequest(ctx context.Context, request interface{}) (pbRequest interface{}, err error) {
	if request == nil {
		return nil, errors.New("nil HelloRequest")
	}
	req := request.(endpoint.HelloRequest)
	pbRequest = &api.HelloRequest{
		Name: req.Name,
	}

	return pbRequest, nil
}

func DecodeHelloRequest(ctx context.Context, pbRequest interface{}) (request interface{}, err error) {
	if pbRequest == nil {
		return nil, errors.New("nil HelloRequest")
	}
	req := pbRequest.(*api.HelloRequest)
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
	pb := &api.HelloResponse{
		Greeting: &api.Greeting{
			Text: resp.Greeting,
		},
	}

	return pb, nil
}

func DecodeHelloResponse(ctx context.Context, pbResponse interface{}) (response interface{}, err error) {
	if pbResponse == nil {
		return nil, errors.New("nil HelloResponse")
	}
	resp := pbResponse.(*api.HelloResponse)
	response = endpoint.HelloResponse{
		Greeting: resp.Greeting.Text,
	}

	return response, nil
}
