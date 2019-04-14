package main

import (
	"strings"

	"github.com/dan-compton/funk/pkg/plugins"
	"github.com/pkg/errors"
)

// ToUpperCaller converts a string to upper case.
type ToUpperCaller struct{}

// Namespace returns the key by which ToUpperCaller can be called in the "go" namespace.
func (e *ToUpperCaller) Namespace() string {
	return "to_upper"
}

// Call decodes and checks the arguments, then calls the ToUpper function.
func (e *ToUpperCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	expectedArgs := plugins.NewArgs(argv0) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	return e.ToUpper(argv0)
}

// ToUpper rewrites a string such that the first letter is capitalized.
func (e *ToUpperCaller) ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}
