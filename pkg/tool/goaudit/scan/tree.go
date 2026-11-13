package scan

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path"
)

func Tree(t *warm_tree.Tree) *virtual_file_system.System {
	result := virtual_file_system.New()
	var visit func(string)
	visit = func(directory string) {
		for _, name := range t.Children(directory) {
			relative := path.Join(directory, name)
			e, _ := t.Stat(relative)

			if e.Directory {
				if !ignoredDirectory(name) &&
					!t.IsFile(path.Join(relative, goModConstant.ModFile)) {
					visit(relative)
				}

				continue
			}

			b, f := t.Read(relative)
			errors.PanicOnError(f)
			result.AddFile(relative, b, e.Modified)
		}
	}
	visit(soil.CurrentDirectory)

	return result
}
