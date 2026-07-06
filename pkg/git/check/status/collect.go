package status

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"strings"
)

func collect(
	root string,
	depth int,
) []*repository.Repository {
	if root == "" {
		root = system.WorkDirectory()
	}

	var result []*repository.Repository
	errors.PanicOnError(
		filepath.Walk(
			root,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if strings.Count(
					system.RelativePath(root, path),
					string(os.PathSeparator),
				) > depth {
					if i.IsDir() {
						return filepath.SkipDir
					}

					return nil
				}

				if i.IsDir() && i.Name() == constant.Directory {
					result = append(
						result,
						git.ReadRepository(filepath.Dir(path)),
					)

					return filepath.SkipDir
				}

				return nil
			},
		),
	)

	for _, r := range result {
		r.Validate()
	}

	return result
}
