package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ReadDirectory(name string) []os.DirEntry {
	result, e := os.ReadDir(name)
	errors.PanicOnError(e)

	return result
}
