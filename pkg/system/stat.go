package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Stat(name string) os.FileInfo {
	result, e := os.Stat(name)
	errors.PanicOnError(e)

	return result
}
