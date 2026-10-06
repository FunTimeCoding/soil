package terminal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (t *Terminal) Exitf(
	format string,
	a ...any,
) {
	_, e := fmt.Fprintf(t.failure, format, a...)
	errors.PanicOnError(e)
	t.Exit(1)
}
