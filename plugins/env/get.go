package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/zenreach/funk/pkg/plugins"
)

// GetCaller provides a Call method to get values from the environment.
type GetCaller struct{}

// Namespace returns the key by which GetCaller can be called in the "go" namespace.
func (e *GetCaller) Namespace() string {
	return "get"
}

// Call decodes and checks the arguments, then calls the Get function.
func (e *GetCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	expectedArgs := plugins.NewArgs(argv0) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	return e.Get(argv0)
}

// Get returns the specified value from the environment.
func (e *GetCaller) Get(s string) (string, error) {
	v := os.Getenv(s)
	return v, nil
}
