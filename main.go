package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"flag"

	"github.com/dan-compton/exile/internal"
	"github.com/dan-compton/exile/internal/plugins/conventions"
	"github.com/dan-compton/exile/internal/plugins/env"
	"github.com/dan-compton/exile/internal/plugins/golang"
	"github.com/dan-compton/exile/internal/plugins/strings"
	"github.com/dan-compton/exile/pkg/plugins"

	"github.com/pkg/errors"
)

const (
	DefaultTemplatePath     = "$GOPATH/src/github.com/dan-compton/exile/examples/"
	DefaultPluginsPath      = "$GOPATH/src/github.com/dan-compton/exile/plugins/"
	DefaultPluginsExtension = ".so"
)

var (
	mappers      []plugins.Mapper
	pluginsPath  string
	outPath      string
	help         bool
	templateRoot string
	extensions   = map[internal.TemplateType]string{
		internal.GlobalInclude: ".t_include",
		internal.Renderable:    ".t",
	}
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.BoolVar(&help, "h", false, "print usage message")
	flag.StringVar(&outPath, "o", pwd, fmt.Sprintln("Output path used when rendering templates.  Default is \"%s\".  Environmental variables are automatically expanded.", pwd))
	flag.StringVar(&templateRoot, "t", DefaultTemplatePath, fmt.Sprintf("The absolute path to the template root/base directory. Default is \"%s\".  Environmental variables are automatically expanded.", DefaultTemplatePath))
	flag.StringVar(&pluginsPath, "p", DefaultPluginsPath, fmt.Sprintf("The absolute path to the plugins root/base directory. Default is \"%s\".  Environmental variables are automatically expanded.", DefaultPluginsPath))
	flag.Parse()

	pluginsPath = os.ExpandEnv(pluginsPath)
	templateRoot = os.ExpandEnv(templateRoot)
}

func main() {
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// XXX hack around lack of darwin support for now.
	switch runtime.GOOS {
	case "linux":
		// Load and open all go plugins from plugins path.
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
	default:
		mappers = append(mappers, env.PluginMappers)
		mappers = append(mappers, conventions.PluginMappers)
		mappers = append(mappers, golang.PluginMappers)
		mappers = append(mappers, strings.PluginMappers)
	}

	// Get a slice of template paths for each template type.
	templatePaths, err := internal.EnumerateTemplates(templateRoot, extensions)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "enumerating templates"))
	}

	// Get a slice of template path -> renderable template
	templates, err := internal.ParseTemplates(templateRoot, templatePaths, extensions, mappers...)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "parsing templates"))
	}

	// render the templates to file
	for relOut, t := range templates {
		ofs := filepath.Join(outPath, relOut)
		log.Println("RENDERING TEMPLATE", ofs)

		op := path.Dir(ofs)
		err := os.MkdirAll(op, os.ModePerm)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "creating folders for template output"))
		}

		ofn := filepath.Base(ofs)
		of, err := os.Create(ofs)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "creating file for template output"))
		}
		defer of.Close()

		err = t.ExecuteTemplate(of, ofn+extensions[internal.Renderable], &struct{}{})
		if err != nil {
			log.Fatalln(errors.Wrap(err, "executing template"))
		}
	}
	log.Println("DONE")
}
