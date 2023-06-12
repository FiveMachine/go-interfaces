package common

import (
	"github.com/taubyte/go-interfaces/vm"
	structureSpec "github.com/taubyte/go-specs/structure"
)

// FunctionContext wraps the methods used for interacting with a WASM function.
type FunctionContext struct {
	Config      structureSpec.Function
	Project     string
	Application string
}

type Function interface {
	// Instantiate creates a Function instance for the given Function
	Instantiate(ctx FunctionContext, branch, commit string) (sdk FunctionInstance, runtime vm.Runtime, sdkPlugin interface{}, err error)
	// Verbose assigns the verbose variable to the Function, used for debugging
	Verbose() bool
}

type FunctionInstance interface {
	// Function returns the parent Function interface
	Function() Function
	// Call will call the Function on the runtime, with an injected parameter of the event id
	Call(runtime vm.Runtime, id interface{}) error
	// Name returns the name of the Function
	Name() string
}
