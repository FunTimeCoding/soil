package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Formatted(format string, a ...any) {
	_, e := fmt.Printf(format, a...)
	errors.PanicOnError(e)
}
