package vm

import "io"

type Backend interface {
	// Returns the URI scheme the backend supports.
	Scheme() string
	// Get attempts to retrieve the WASM asset.
	Get(uri string) (io.ReadCloser, error)
	// Close will close the Backend.
	Close() error
}

type Resolver interface {
	// Lookup resolves a module name and returns the uri
	Lookup(ctx Context, module string) (string, error)
}

type Loader interface {
	// Load resolves the module, then loads the module using a Backend
	Load(ctx Context, module string) (io.ReadCloser, error)
}

type Source interface {
	// Module Loads the given module name, and returns the SourceModule
	Module(ctx Context, name string) (SourceModule, error)
}

type SourceModule interface {
	// Source returns the raw data of the source
	Source() []byte

	// Imports returns functions, and memories required for instantiation
	Imports() []string

	// Imports returns functions, and memories for the specific module.
	ImportsByModule(name string) []string

	// ImportsFunction returns a boolean based on existence of a function in given module.
	ImportsFunction(module, name string) bool
}
