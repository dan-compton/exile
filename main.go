package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"flag"

	"github.com/dan-compton/exile/internal"
	"github.com/dan-compton/exile/pkg/plugins"

	"github.com/pkg/errors"
)

const (
	DefaultTemplatePath      = "$GOPATH/src/github.com/dan-compton/exile/examples/"
	DefaultTemplateExtension = ".t"
	DefaultPluginsPath       = "$GOPATH/src/github.com/dan-compton/exile/plugins/"
	DefaultPluginsExtension  = ".so"
)

var (
	mappers      []plugins.Mapper
	templatePath string
	pluginsPath  string
	packagePath  string
	outPath      string
	help         bool
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.BoolVar(&help, "h", false, "print usage message")
	flag.StringVar(&outPath, "o", pwd, fmt.Sprintln("Output path used when rendering templates.  Default is \"%s\".  Environmental variables are automatically expanded.", pwd))
	flag.StringVar(&templatePath, "t", DefaultTemplatePath, fmt.Sprintf("The absolute path to the template root/base directory. Default is \"%s\".  Environmental variables are automatically expanded.", DefaultTemplatePath))
	flag.StringVar(&pluginsPath, "p", DefaultPluginsPath, fmt.Sprintf("The absolute path to the plugins root/base directory. Default is \"%s\".  Environmental variables are automatically expanded.", DefaultPluginsPath))
	flag.Parse()

	// expand environmental variables for inputs.
	templatePath = os.ExpandEnv(templatePath)
	packagePath = os.ExpandEnv(packagePath)
	pluginsPath = os.ExpandEnv(pluginsPath)
}

func main() {
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	templates, err := internal.EnumerateTemplates(templatePath, DefaultTemplateExtension)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "enumerating templates in \"%s\"", templatePath))
	}

	plugs, err := internal.LoadPlugins(pluginsPath, DefaultPluginsExtension)
	if err != nil {
		log.Fatalln(errors.Wrapf(err, "enumerating plugins in \"%s\"", pluginsPath))
	}
	for _, plug := range plugs {
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "loading plugin"))
		}
		m, err := plug.Lookup("PluginMappers")
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "looking up plugin symbol PluginMappers"))
		}

		mappers = append(mappers, m.(plugins.Mapper))
	}

	for _, t := range templates {
		t, err := t.Parse(mappers...)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("RENDERING TEMPLATE", filepath.Join(outPath, t.RelativeOutputPath()))
		err = t.Render(outPath)
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("DONE")
}
