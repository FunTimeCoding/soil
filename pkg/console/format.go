package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Format(format string, a ...any) {
	_, e := fmt.Printf(format, a...)
	errors.PanicOnError(e)
}
