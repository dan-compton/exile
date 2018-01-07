package main

import (
	"strings"

	"github.com/dan-compton/exile/pkg/plugins"
	"github.com/pkg/errors"
)

type UnexportedCaller struct{}

// Namespace returns the key by which UnexportedCaller can be called in the "go" namespace.
func (e *UnexportedCaller) Namespace() string {
	return "exported"
}

// Call decodes and checks the arguments, then calls the Unexported function.
func (e *UnexportedCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	expectedArgs := plugins.NewArgs(argv0) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	return e.Unexported(argv0)
}

// Unexported rewrites a string such that the first letter is capitalized.
func (e *UnexportedCaller) Unexported(s string) (string, error) {
	return strings.ToLower(string(s)[:1]) + string(s)[1:], nil
}
