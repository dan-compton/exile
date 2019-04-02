package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/dan-compton/funk/pkg/plugins"
)

// SetCaller provides a Call method to set values from the environment.
type SetCaller struct{}

// Namespace returns the key by which SetCaller can be called in the "go" namespace.
func (e *SetCaller) Namespace() string {
	return "set"
}

// Call decodes and checks the arguments, then calls the Set function.
func (e *SetCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	argv1 := "string"
	expectedArgs := plugins.NewArgs(argv0, argv1) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	receivedArgs[1].Set(&argv1)
	return e.Set(argv0, argv1)
}

// Set sets the specified value in the environment.
func (e *SetCaller) Set(k, v string) (string, error) {
	os.Setenv(k, v)
	return "", nil
}
