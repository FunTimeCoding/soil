package resolve

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/system"
)

func ExitOnLoadFailure(e error) {
	if f, okay := load_failure.As(e); okay {
		system.Exitf(1, "does not load: %s - %v\n", f.Package, f.Cause)
	}

	errors.PanicOnError(e)
}
