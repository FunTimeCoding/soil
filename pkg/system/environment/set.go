package environment

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func Set(
	k string,
	v string,
) {
	errors.PanicOnError(os.Setenv(k, v))
}
