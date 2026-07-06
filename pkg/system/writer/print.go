package writer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func Print(
	w io.Writer,
	format string,
	a ...any,
) {
	_, e := fmt.Fprintf(w, format, a...)
	errors.PanicOnError(e)
}
