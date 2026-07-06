package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func ExecutablePath() string {
	result, e := os.Executable()
	errors.PanicOnError(e)

	return result
}
