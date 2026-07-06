package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func MakeDirectory(path string) {
	errors.PanicOnError(os.MkdirAll(path, 0755))
}
