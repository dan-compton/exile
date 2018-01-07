package internal

import (
	"os"
	"path/filepath"
	"plugin"

	"github.com/pkg/errors"
)

// LoadPlugins returns a plugin for each found plugin in the path.
func LoadPlugins(root, ext string) ([]*plugin.Plugin, error) {
	var plugins []*plugin.Plugin
	err := filepath.Walk(root, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(currPath) == ext {
			p, err := plugin.Open(currPath)
			if err != nil {
				return errors.Wrap(err, "enumerating plugins")
			}
			plugins = append(plugins, p)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "enumerating plugins in target layout directory")
	}
	return plugins, nil
}
