package system

import (
	"fmt"
	"os"
)

func Exitf(
	code int,
	format string,
	a ...any,
) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(code)
}
