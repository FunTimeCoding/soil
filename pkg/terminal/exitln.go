package terminal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (t *Terminal) Exitln(a ...any) {
	_, e := fmt.Fprintln(t.failure, a...)
	errors.PanicOnError(e)
	t.Exit(1)
}
