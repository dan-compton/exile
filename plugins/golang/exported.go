package main

import (
	"strings"

	"github.com/dan-compton/funk/pkg/plugins"
	"github.com/pkg/errors"
)

type ExportedCaller struct{}

// Namespace returns the key by which ExportedCaller can be called in the "go" namespace.
func (e *ExportedCaller) Namespace() string {
	return "exported"
}

// Call decodes and checks the arguments, then calls the Exported function.
func (e *ExportedCaller) Call(is ...interface{}) (string, error) {
	argv0 := "string"
	expectedArgs := plugins.NewArgs(argv0) // only used for type equality.
	receivedArgs := plugins.NewArgs(is...)

	// check arg position and type equality.
	if !expectedArgs.Equals(receivedArgs) {
		return "", errors.Errorf("%s expects arguments of the form %s", e.Namespace(), expectedArgs.String())
	}

	// set expected args by position.
	receivedArgs[0].Set(&argv0)
	return e.Exported(argv0)
}

// Exported rewrites a string such that the first letter is capitalized.
func (e *ExportedCaller) Exported(s string) (string, error) {
	return strings.ToUpper(string(s)[:1]) + string(s)[1:], nil
}
