package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func TemporaryDirectory(pattern string) string {
	result, e := os.MkdirTemp("", pattern)
	errors.PanicOnError(e)

	return result
}
