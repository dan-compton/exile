package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/zenreach/exile/pkg/plugins"
)

type ToLowerCaller struct{}

// Namespace returns the key by which ToLowerCaller can be called in the "go" namespace.
func (e *ToLowerCaller) Namespace() string {
	return "to_lower"
}

// Call decodes and checks the arguments, then calls the ToLower function.
func (e *ToLowerCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	expectedArgs := plugins.NewArgs(argv0) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	return e.ToLower(argv0)
}

// ToLower rewrites a string such that the first letter is capitalized.
func (e *ToLowerCaller) ToLower(s string) (string, error) {
	return strings.ToLower(s), nil
}
