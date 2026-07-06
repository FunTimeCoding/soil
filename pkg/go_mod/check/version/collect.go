package version

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/go_mod/check/version/option"
	"github.com/funtimecoding/soil/pkg/go_mod/project"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"strings"
)

func collect(o *option.Version) []*project.Project {
	if o.Path == "" {
		o.Path = system.WorkDirectory()
	}

	var result []*project.Project
	errors.PanicOnError(
		filepath.Walk(
			o.Path,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if strings.Count(
					system.RelativePath(o.Path, path),
					string(os.PathSeparator),
				) > o.Depth {
					if i.IsDir() {
						return filepath.SkipDir
					}

					return nil
				}

				if exclude(i.Name(), o.Skip) {
					if i.IsDir() {
						return filepath.SkipDir
					}
				}

				if !i.IsDir() && i.Name() == go_mod.ModFile {
					result = append(
						result,
						go_mod.ReadProject(
							filepath.Dir(path),
							o.RuntimeVersion,
						),
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
