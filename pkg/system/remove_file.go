package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func RemoveFile(path string) {
	errors.PanicOnError(os.Remove(path))
}
