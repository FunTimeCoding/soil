package build_tag

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path/filepath"
	"strings"
)

func Scanned(
	directory string,
	path string,
) bool {
	if !strings.HasSuffix(path, constant.GoExtension) {
		return false
	}

	relative, e := filepath.Rel(directory, filepath.Dir(path))

	if e != nil || !filepath.IsLocal(relative) {
		return false
	}

	if relative == constant.CurrentDirectory {
		return true
	}

	for _, name := range strings.Split(relative, string(filepath.Separator)) {
		if skipped(name) {
			return false
		}
	}

	return true
}
