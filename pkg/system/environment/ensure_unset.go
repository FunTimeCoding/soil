package environment

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func EnsureUnset(name string) {
	if os.Getenv(name) != "" {
		errors.PanicOnError(os.Unsetenv(name))
	}
}
