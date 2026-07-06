package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Home() string {
	result, e := os.UserHomeDir()
	errors.PanicOnError(e)

	return result
}
