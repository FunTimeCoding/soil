package repository

import (
	"github.com/funtimecoding/soil/pkg/constant"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path"
	"path/filepath"
)

func siblingCheckouts(
	root string,
	modules []string,
) []string {
	var result []string

	for _, m := range modules {
		name := path.Base(m)

		if lintConstant.MajorSuffix.MatchString(name) {
			name = path.Base(path.Dir(m))
		}

		sibling := filepath.Join(constant.ParentDirectory, name)

		if system.DirectoryExists(filepath.Join(root, sibling)) {
			result = append(result, sibling)
		}
	}

	return result
}
