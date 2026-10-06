package terminal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (t *Terminal) Line(a ...any) {
	_, e := fmt.Fprintln(t.output, a...)
	errors.PanicOnError(e)
}
