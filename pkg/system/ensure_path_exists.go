package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func EnsurePathExists(path string) {
	if _, e := os.Stat(path); os.IsNotExist(e) {
		errors.LogOnError(os.MkdirAll(path, 0755))
	}
}
