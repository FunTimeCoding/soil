package terminal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
)

func (t *Terminal) Blockln(
	code int,
	a ...any,
) {
	_, e := fmt.Fprintln(t.failure, a...)
	errors.PanicOnError(e)
	t.end(constant.Blocked, code)
}
