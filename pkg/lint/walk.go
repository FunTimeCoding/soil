package lint

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path"
	"path/filepath"
)

func Walk(
	root string,
	o *option.Lint,
) (*repository.Repository, []string) {
	absolute := system.AbsolutePath(root)
	t := o.Trees.Tree(absolute)
	v := virtual_file_system.New()
	var empty []string
	var nested []string
	directories := map[string]bool{soil.CurrentDirectory: true}
	var visit func(string)
	visit = func(directory string) {
		for _, name := range t.Children(directory) {
			relative := path.Join(directory, name)
			e, _ := t.Stat(relative)

			if e.Directory {
				if e.Linked {
					directories[relative] = true

					continue
				}

				if t.IsFile(path.Join(relative, goModConstant.ModFile)) {
					nested = append(nested, relative)

					continue
				}

				directories[relative] = true

				if !Skipped(
					o.Skips,
					join.Empty(relative, stringsConstant.Slash),
				) &&
					len(t.Children(relative)) == 0 &&
					system.IsEmptyDirectory(filepath.Join(absolute, relative)) {
					empty = append(empty, relative)
				}

				visit(relative)

				continue
			}

			if o.Metadata || Skipped(
				o.Skips,
				join.Empty(path.Dir(relative), stringsConstant.Slash),
			) {
				v.AddMetadata(relative, e.Size, e.Modified)

				continue
			}

			b, f := t.Read(relative)
			errors.PanicOnError(f)
			v.AddFile(relative, b, e.Modified)
		}
	}
	visit(soil.CurrentDirectory)
	result := repository.NewWithForest(root, v, o.Trees)
	result.NestedModules = nested
	result.Directories = directories

	return result, empty
}
