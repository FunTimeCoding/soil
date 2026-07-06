package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Remove(path string) {
	errors.PanicOnError(os.RemoveAll(path))
}
