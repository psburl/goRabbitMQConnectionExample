package environment

import "os"

// Manager aims to manage environment configuration of code.
type Manager interface {
	GetAmqpURL() string
}

// Production environment implements environment management of variables to production release
type Production struct {
}

// NewProduction constructs a new env manager for production behavior
func NewProduction() *Production {
	return &Production{}
}

// GetAmqpURL returns the value of "AMQP_URL" environment variable
func (*Production) GetAmqpURL() string {
	return os.Getenv("AMQP_URL")
}
