package plugins

import (
	"text/template"
)

// Mapper is a thing that maps string transformation functions to a template.FuncMap.
// NOTE: A method or function must conform to the template.FuncMap spec.
type Mapper interface {
	Map(t template.FuncMap)
}
