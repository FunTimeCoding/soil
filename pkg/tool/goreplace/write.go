package goreplace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func write(
	out io.Writer,
	format string,
	a ...any,
) {
	_, e := fmt.Fprintln(out, fmt.Sprintf(format, a...))
	errors.PanicOnError(e)
}
