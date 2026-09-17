package lint

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"strings"
)

func Scopes(
	root string,
	work string,
	paths []string,
) ([]string, error) {
	var result []string

	for _, p := range paths {
		absolute := p

		if !filepath.IsAbs(p) {
			absolute = filepath.Join(work, p)
		}

		relative, e := filepath.Rel(root, absolute)

		if e != nil {
			return nil, e
		}

		relative = filepath.ToSlash(relative)

		if relative == constant.ParentDirectory ||
			strings.HasPrefix(relative, "../") {
			return nil, errors.Format("path outside repository", p)
		}

		if !system.FileExists(absolute) &&
			!system.DirectoryExists(absolute) {
			return nil, errors.Format("path does not exist", p)
		}

		if relative == constant.CurrentDirectory {
			continue
		}

		result = append(result, relative)
	}

	return result, nil
}
