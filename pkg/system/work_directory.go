package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func WorkDirectory() string {
	d, e := os.Getwd()
	errors.PanicOnError(e)

	return d
}
