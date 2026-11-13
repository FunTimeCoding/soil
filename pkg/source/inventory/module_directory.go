package inventory

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func (i *Inventory) ModuleDirectory(modulePath string) (string, bool) {
	for _, m := range i.Modules {
		content, e := os.ReadFile(
			filepath.Join(m.Directory, goModConstant.ModFile),
		)

		if e == nil && modfile.ModulePath(content) == modulePath {
			return m.Directory, true
		}
	}

	return "", false
}
