package system

import (
	"errors"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"os/exec"
)

func ExitErrorCode(e error) int {
	if e != nil {
		if f, okay := errors.AsType[*exec.ExitError](e); okay {
			return f.ExitCode()
		}

		soilErrors.PanicOnError(e)
	}

	return 0
}
