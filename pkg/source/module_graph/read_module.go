package module_graph

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func readModule(root string) (string, map[string]string, map[string]string) {
	p := filepath.Join(root, constant.ModFile)
	b, e := os.ReadFile(p)
	errors.PanicOnError(e)
	f, e := modfile.Parse(p, b, nil)
	errors.PanicOnError(e)
	requirements := make(map[string]string, len(f.Require))

	for _, r := range f.Require {
		requirements[r.Mod.Path] = r.Mod.Version
	}

	replaced := make(map[string]string)

	for _, r := range f.Replace {
		if !modfile.IsDirectoryPath(r.New.Path) {
			continue
		}

		directory := r.New.Path

		if !filepath.IsAbs(directory) {
			directory = filepath.Join(root, directory)
		}

		replaced[r.Old.Path] = directory
	}

	return f.Module.Mod.Path, requirements, replaced
}
