package console

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Line(a ...any) {
	_, e := fmt.Println(a...)
	errors.PanicOnError(e)
}
