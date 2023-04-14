package vm

import "context"

type Context interface {
	// Context returns the go context of the function instance
	Context() context.Context

	// Close calls the go context cancel method.
	Close() error

	// Project returns the Taubyte project id
	Project() string

	// Application returns the application, if none returns an empty string
	Application() string

	// Resource returns the id of the resource being used.
	Resource() string

	// Branch returns the branch name used by this resource execution pipeline.
	Branch() string

	// Commit returns the commit id used by this resource execution pipeline.
	Commit() string
}
