package lint

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func Walk(
	root string,
	o *option.Lint,
) (*repository.Repository, []string) {
	v := virtual_file_system.New()
	var empty []string
	errors.PanicOnError(
		filepath.WalkDir(
			root,
			func(
				p string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				relative, f := filepath.Rel(root, p)
				errors.PanicOnError(f)
				relative = filepath.ToSlash(relative)

				if relative == constant.CurrentDirectory {
					return nil
				}

				if d.IsDir() {
					if d.Name() == gitConstant.Directory {
						return filepath.SkipDir
					}

					if !Skipped(o, join.Empty(relative, stringsConstant.Slash)) &&
						system.IsEmptyDirectory(p) {
						empty = append(empty, relative)
					}

					return nil
				}

				if d.Type()&fs.ModeSymlink != 0 {
					return nil
				}

				i, g := d.Info()
				errors.PanicOnError(g)

				if Skipped(o, join.Empty(path.Dir(relative), stringsConstant.Slash)) {
					v.AddMetadata(relative, i.Size(), i.ModTime())

					return nil
				}

				b, h := os.ReadFile(p)
				errors.PanicOnError(h)
				v.AddFile(relative, b, i.ModTime())

				return nil
			},
		),
	)

	return repository.New(root, v), empty
}
