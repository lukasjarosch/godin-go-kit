// This is the core service file. It specifies the behaviour of the business-logic
// and enables us to easily layer our application.
package service

import (
	"context"
)

// ExampleService defines the exposed business logic of the application.
// Each function of the ExampleService is exposed via endpoints.
type ExampleService interface {
	Hello(ctx context.Context, name string) (greeting string, err error)
}

// This is the repository which the ExampleService requires for data access
type Repository interface {
}

// Here shall be space for the internal data types

