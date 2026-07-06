package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Create(name string) *os.File {
	result, e := os.Create(name)
	errors.PanicOnError(e)

	return result
}
