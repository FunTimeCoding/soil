package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"path/filepath"
)

func Match(
	pattern string,
	name string,
) bool {
	result, e := filepath.Match(pattern, name)
	errors.PanicOnError(e)

	return result
}
