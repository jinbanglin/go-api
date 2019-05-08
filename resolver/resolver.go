// Package resolver resolves a http request to an endpoint
package resolver

import (
	"net/http"
)

// Resolver resolves requests to endpoints
type Resolver interface {
	Resolve(r *http.Request) (*Endpoint, error)
	String() string
}

// Endpoint is the endpoint for a http request
type Endpoint struct {
	// e.g greeter
	Name string
	// HTTP Host e.g example.com
	Host string
	// HTTP Methods e.g GET, POST
	Method string
	// HTTP Path e.g /greeter.
	Path string
}
