package main

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/zenreach/exile/pkg/plugins"
)

// SnakeCaller provides a Call method to set values from the environment.
type SnakeCaller struct{}

// Namespace returns the key by which SnakeCaller can be called in the "go" namespace.
func (e *SnakeCaller) Namespace() string {
	return "snake"
}

// Call decodes and checks the arguments, then calls the Set function.
func (e *SnakeCaller) Call(is ...interface{}) (string, error) {
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
	return e.Snake(argv0, argv1)
}

// Snake converts the given string `s` to snake case given the convention `from`.
func (e *SnakeCaller) Snake(from string, s string) (string, error) {
	switch from {
	case "from_camel":
		var firstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		var otherCap = regexp.MustCompile("([a-z0-9])([A-Z])")

		s = firstCap.ReplaceAllString(s, "${1}_${2}")
		s = otherCap.ReplaceAllString(s, "${1}_${2}")
	default:
		return "", errors.Errorf("unknown convention %s", from)
	}
	return s, nil
}
