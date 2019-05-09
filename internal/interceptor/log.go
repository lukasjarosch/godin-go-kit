package interceptor

import (
	"github.com/go-kit/kit/log"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
	"context"
	"time"
)

type logInterceptor struct {
	logger log.Logger
	next service.ExampleService
}

func NewLogInterceptor(logger log.Logger) Interceptor {
	return func(next service.ExampleService) service.ExampleService {
		return &logInterceptor{logger, next}
	}
}

func (i logInterceptor) Hello(ctx context.Context, name string) (greeting string, err error) {
	i.logger.Log("endpoint", "Hello", "request", struct {
		Name string
	}{Name:name})

	// TODO: request-id extraction from context
	// TODO: replace structs with internal req/resp structures

	defer func(begin time.Time) {
		i.logger.Log("endpoint", "Hello", "response", struct {
			Greeting string
		}{Greeting:greeting}, "took", time.Since(begin))
	}(time.Now())

	return i.next.Hello(ctx, name)
}