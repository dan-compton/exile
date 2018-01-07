package internal

import (
	"github.com/dan-compton/exile/pkg/plugins"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// NewTemplate returns a new template from the `root` directory.
func NewTemplate(root, templatePath, ext string) (*Template, error) {
	// given root path /templates/root and template path /templates/root/subdir/template.t
	// determine relative path to template by removing root prefix from template path.
	// => ./subdir/template.t
	rp := strings.TrimPrefix(filepath.Clean(templatePath), filepath.Clean(root))
	rp = path.Dir(rp)

	// determine template filename by removing template suffix.
	of := filepath.Base(strings.TrimSuffix(filepath.Base(templatePath), ext))
	if of == "" {
		return nil, errors.New("template file has no output filename")
	}

	// join the relative path and output filename to get complete relative output.
	ro := path.Join(rp, of)
	if len(ro) > 0 {
		ro = ro[1:]
	}

	return &Template{
		Template: template.New(filepath.Base(templatePath)),
		path:     templatePath,
		of:       of,
		rp:       rp,
		ro:       ro,
	}, nil
}

// EnumerateTemplates recursively parses files with `ext` starting at `root`,
// yielding all desired templates, or error.
// NOTE: Parse must still be called on returned templates.
func EnumerateTemplates(root, ext string) ([]*Template, error) {
	var templates []*Template
	err := filepath.Walk(root, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(currPath) == ext {
			t, err := NewTemplate(root, currPath, ext)
			if err != nil {
				return errors.Wrap(err, "enumerating templates")
			}
			templates = append(templates, t)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "enumerating templates in target layout directory")
	}
	return templates, nil
}

// Template wraps a go template, providing various helpers used for rendering.
type Template struct {
	*template.Template
	path string
	of   string // outpuf filename
	rp   string // relative package
	ro   string // relative output path
}

// OutputFileName returns the output file name.
func (t *Template) OutputFileName() string {
	return t.of
}

// RelativeOutputPath returns the relative output path.
func (t *Template) RelativeOutputPath() string {
	return t.ro
}

// Path returns the output path.
func (t *Template) Path() string {
	return t.path
}

// Parse parses the supplied template path.
// NOTE: this must be called after Funcs.
func (t *Template) Parse(fms ...plugins.Mapper) (*Template, error) {
	var newTemplate *template.Template
	tfs := make(template.FuncMap)
	for _, fm := range fms {
		fm.Map(tfs)
	}
	newTemplate = t.Template.Funcs(tfs)

	pt, err := newTemplate.ParseFiles(t.path)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing template at path %s", t.path)
	}
	newTemplate = pt
	return &Template{
		newTemplate,
		t.path,
		t.of,
		t.rp,
		t.ro,
	}, nil
}

// Render writes the template to a file at path absPackage + t.relOutput.
func (t *Template) Render(absPackage string, data ...interface{}) error {
	outFile := filepath.Join(absPackage, t.ro)
	outPath := path.Dir(outFile)
	err := os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "creating folders for template output")
	}
	of, err := os.Create(outFile)
	if err != nil {
		return errors.Wrap(err, "creating file for template output")
	}
	defer of.Close()
	err = t.Execute(of, data)
	if err != nil {
		return errors.Wrap(err, "executing template")
	}
	return nil
}
