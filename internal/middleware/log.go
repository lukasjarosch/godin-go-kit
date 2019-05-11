package middleware

import (
	"github.com/go-kit/kit/log"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
	"context"
	"time"
)

type logMiddleware struct {
	logger log.Logger
	next   example.ExampleService
}

func NewLogMiddleware(logger log.Logger) Middleware {
	return func(next example.ExampleService) example.ExampleService {
		return &logMiddleware{logger, next}
	}
}

func (i logMiddleware) Hello(ctx context.Context, name string) (greeting string, err error) {
	i.logger.Log("endpoint", "Hello", "request", struct {
		Name string
	}{Name:name})

	// TODO: request-id extraction from context
	// TODO: replace structs with internal req/resp structures

	defer func(begin time.Time) {
		if err != nil {
			i.logger.Log("endpoint", "Hello", "response", struct {
				Greeting string
			}{Greeting:greeting}, "took", time.Since(begin),
				"err", err.Error())
		}

		i.logger.Log("endpoint", "Hello", "response", struct {
			Greeting string
		}{Greeting:greeting}, "took", time.Since(begin))
	}(time.Now())

	return i.next.Hello(ctx, name)
}