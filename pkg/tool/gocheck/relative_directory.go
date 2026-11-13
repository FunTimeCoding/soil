package gocheck

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"path/filepath"
)

func relativeDirectory(
	root string,
	directory string,
) string {
	result, e := filepath.Rel(root, directory)
	errors.PanicOnError(e)

	return filepath.ToSlash(result)
}
