package resolve

import (
	"github.com/funtimecoding/soil/pkg/errors"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func LocalReplacements(directory string) []string {
	p := filepath.Join(directory, go_modConstant.ModFile)

	if !system.FileExists(p) {
		return nil
	}

	b, e := os.ReadFile(p)
	errors.PanicOnError(e)
	f, e := modfile.Parse(p, b, nil)
	errors.PanicOnError(e)
	var result []string

	for _, r := range f.Replace {
		if modfile.IsDirectoryPath(r.New.Path) {
			result = append(result, r.Old.Path)
		}
	}

	return result
}
