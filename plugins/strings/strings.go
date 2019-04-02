package main

import (
	"github.com/dan-compton/funk/pkg/plugins"
)

// PluginNamespace is the function namespace used for the plugin call function.
// It is the same as the filename by convention.
const PluginNamespace = "strings"

// PluginMappers is used to register functions in template.FuncMap.
var PluginMappers = plugins.NewMappers(PluginNamespace)

func init() {
	PluginMappers.Register(new(ToUpperCaller), new(ToLowerCaller))
}
