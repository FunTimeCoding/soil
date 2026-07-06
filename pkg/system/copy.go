package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func Copy(
	source io.Reader,
	destination io.Writer,
) {
	_, e := io.Copy(destination, source)
	errors.PanicOnError(e)
}
