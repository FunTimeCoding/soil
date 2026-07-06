package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func DeleteFile(name string) {
	errors.PanicOnError(os.Remove(name))
}
