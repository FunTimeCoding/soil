package build_tag

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func Discover(directory string) []string {
	if directory == "" {
		directory = constant.CurrentDirectory
	}

	tags := make(map[string]bool)
	errors.PanicOnError(
		filepath.WalkDir(
			directory,
			func(
				path string,
				n fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return nil
				}

				if n.IsDir() {
					name := n.Name()

					if name == "vendor" ||
						name == "testdata" ||
						name == "tmp" ||
						name != constant.CurrentDirectory &&
							strings.HasPrefix(
								name,
								constant.CurrentDirectory,
							) {
						return filepath.SkipDir
					}

					return nil
				}

				if !strings.HasSuffix(path, constant.GoExtension) {
					return nil
				}

				for _, t := range extract(path) {
					tags[t] = true
				}

				return nil
			},
		),
	)
	result := make([]string, 0, len(tags))

	for t := range tags {
		result = append(result, t)
	}

	sort.Strings(result)

	return result
}
