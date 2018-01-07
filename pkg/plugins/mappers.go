package plugins

import (
	"text/template"

	"github.com/pkg/errors"
)

// NewMappers returns a new Mappers of provided namespace.
func NewMappers(namespace string) Mappers {
	return Mappers{
		callers:   make(map[string]Caller),
		namespace: namespace,
	}
}

// Mappers is a map of caller functions.
// NOTE: Mappers is not safe for concurrent use.
type Mappers struct {
	callers   map[string]Caller
	namespace string
}

// Namespace returns the key at which Call function should be mapped.
func (m *Mappers) Namespace() string {
	return m.namespace
}

// Call calls the function keyed by key.
func (m *Mappers) Call(inputs ...interface{}) (string, error) {
	var namespace string
	// consume the first input as namespace
	if len(inputs) >= 1 {
		switch v := inputs[0].(type) {
		case string:
			namespace = v
		default:
			return "", errors.Errorf("caller expected type %T, got %T", new(string), namespace)
		}
	}
	// fetch and call caller at namespace
	caller, ok := m.callers[namespace]
	if !ok {
		return "", errors.Errorf("function %s not found", namespace)
	}
	inputs = inputs[1:]
	return caller.Call(inputs...)
}

// Register registers a new caller function.
func (m *Mappers) Register(cs ...Caller) {
	for _, c := range cs {
		m.callers[c.Namespace()] = c
	}
}

// Map maps the call function into the appropriate namespace.
func (m *Mappers) Map(t template.FuncMap) {
	t[m.namespace] = m.Call
}
