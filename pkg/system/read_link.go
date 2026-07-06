package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ReadLink(name string) string {
	result, e := os.Readlink(name)
	errors.PanicOnError(e)

	return result
}
