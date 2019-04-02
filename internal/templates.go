package internal

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"text/template"

	"github.com/dan-compton/funk/pkg/plugins"
	"github.com/pkg/errors"
)

type TemplateType = int

const (
	GlobalInclude TemplateType = iota
	Renderable
)

// EnumerateTemplates returns a mapping from template type to []string of paths.
func EnumerateTemplates(root string, extensions map[TemplateType]string) (map[TemplateType][]string, error) {
	templates := make(map[TemplateType][]string)
	for t, _ := range extensions {
		_, ok := templates[t]
		if !ok {
			templates[t] = []string{}
		}
	}

	err := filepath.Walk(root, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(currentPath)
		for k, v := range extensions {
			if ext != v {
				continue
			}
			templates[k] = append(templates[k], currentPath)
		}
		return nil
	})
	return templates, err
}

// ParseTemplates returns a map of relative output paths to templates given the templates `root` path and map of `extensions` to
// templates types.
func ParseTemplates(root string, templatePaths map[TemplateType][]string, extensions map[TemplateType]string, mappers ...plugins.Mapper) (map[string]*template.Template, error) {
	globals := templatePaths[GlobalInclude]

	ext, ok := extensions[Renderable]
	if !ok && len(templatePaths[Renderable]) == 0 {
		return nil, errors.New("no extension was provided for renderable templates")
	}

	renderables := make(map[string]*template.Template)
	renderablePaths := templatePaths[Renderable]
	for _, renderablePath := range renderablePaths {
		// given root path /templates/root and template path /templates/root/subdir/template.t
		// determine relative path to template by removing root prefix from template path.
		// => ./subdir/template.t
		rp := strings.TrimPrefix(filepath.Clean(renderablePath), filepath.Clean(root))
		rp = path.Dir(rp)

		// determine template filename by removing template suffix.
		of := filepath.Base(strings.TrimSuffix(filepath.Base(renderablePath), ext))
		if of == "" {
			return nil, errors.Errorf("template file %s has no output filename")
		}

		// join the relative path and output filename to get complete relative output.
		ro := path.Join(rp, of)
		if len(ro) > 0 {
			ro = ro[1:]
		}

		var newTemplate *template.Template
		tfm := make(template.FuncMap)
		for _, mapper := range mappers {
			mapper.Map(tfm)
		}
		newTemplate = template.New(of + ext).Funcs(tfm)

		fs := []string{renderablePath}
		for _, global := range globals {
			fs = append(fs, global)
		}

		templ, err := newTemplate.ParseFiles(fs...)
		if err != nil {
			return nil, errors.Wrap(err, "parsing template")
		}

		renderables[ro] = templ
	}
	return renderables, nil
}
