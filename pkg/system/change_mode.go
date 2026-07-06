package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ChangeMode(
	path string,
	mode os.FileMode,
) {
	errors.PanicOnError(os.Chmod(path, mode))
}
