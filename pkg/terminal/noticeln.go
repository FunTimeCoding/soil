package terminal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (t *Terminal) Noticeln(a ...any) {
	_, e := fmt.Fprintln(t.failure, a...)
	errors.PanicOnError(e)
}
