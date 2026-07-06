package system

import (
	"errors"
	errorLibrary "github.com/funtimecoding/soil/pkg/errors"
	"os/exec"
)

func ExitErrorCode(e error) int {
	if e != nil {
		if f, okay := errors.AsType[*exec.ExitError](e); okay {
			return f.ExitCode()
		}

		errorLibrary.PanicOnError(e)
	}

	return 0
}
