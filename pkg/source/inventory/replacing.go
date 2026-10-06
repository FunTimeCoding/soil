package inventory

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func (i *Inventory) Replacing(directory string) []string {
	active, e := os.ReadFile(filepath.Join(directory, constant.ModFile))

	if e != nil {
		return nil
	}

	modulePath := modfile.ModulePath(active)
	root := filepath.Clean(directory)
	var result []string

	for _, m := range i.Modules {
		other := filepath.Clean(m.Directory)

		if other == root {
			continue
		}

		content, f := os.ReadFile(filepath.Join(other, constant.ModFile))

		if f != nil {
			continue
		}

		parsed, g := modfile.Parse(constant.ModFile, content, nil)

		if g != nil {
			continue
		}

		for _, r := range parsed.Replace {
			if r.Old.Path != modulePath || r.New.Version != "" {
				continue
			}

			target := r.New.Path

			if !filepath.IsAbs(target) {
				target = filepath.Join(other, target)
			}

			if filepath.Clean(target) == root {
				result = append(result, other)
			}
		}
	}

	return result
}
