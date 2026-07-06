package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Hostname() string {
	result, e := os.Hostname()
	errors.PanicOnError(e)

	return result
}
