package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"path/filepath"
)

func AbsolutePath(p string) string {
	result, e := filepath.Abs(p)
	errors.PanicOnError(e)

	return result
}
