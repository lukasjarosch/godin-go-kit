package middleware

import (
	"github.com/lukasjarosch/godin-go-kit/internal/example"
)

type Middleware func(service example.ExampleService) example.ExampleService