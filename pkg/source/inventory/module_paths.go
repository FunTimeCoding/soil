package inventory

import (
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func (i *Inventory) ModulePaths() []string {
	var result []string

	for _, m := range i.Modules {
		content, e := os.ReadFile(
			filepath.Join(m.Directory, go_modConstant.ModFile),
		)

		if e != nil {
			continue
		}

		if p := modfile.ModulePath(content); p != "" {
			result = append(result, p)
		}
	}

	return result
}
