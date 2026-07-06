package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ChangeDirectory(path string) {
	errors.PanicOnError(os.Chdir(path))
}
