package conventions

import (
	"bytes"
	"unicode"

	"github.com/dan-compton/exile/pkg/plugins"
	"github.com/pkg/errors"
)

// CamelCaller provides a Call method to set values from the environment.
type CamelCaller struct{}

// Namespace returns the key by which CamelCaller can be called in the "go" namespace.
func (e *CamelCaller) Namespace() string {
	return "camel"
}

// Call decodes and checks the arguments, then calls the Set function.
func (e *CamelCaller) Call(is ...interface{}) (string, error) {
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
func (e *CamelCaller) Snake(from string, s string) (string, error) {
	switch from {
	case "from_snake":
		if len(s) == 0 {
			return "", nil
		}
		var b bytes.Buffer
		var prev rune
		for i, v := range s {
			if unicode.IsUpper(v) && unicode.IsUpper(prev) {
				b.WriteRune(unicode.ToLower(v))
			} else {
				if unicode.IsUpper(prev) {
					if i != 0 {
						b.Truncate(i - 1)
						b.WriteRune(unicode.ToUpper(prev))
					}
				}
				b.WriteRune(v)
			}
			prev = v
		}
		return s, nil
	default:
		return "", errors.Errorf("unknown convention %s", from)
	}
	return s, nil
}
