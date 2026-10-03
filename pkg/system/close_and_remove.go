package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func closeAndRemove(
	c io.Closer,
	path string,
) {
	errors.PanicClose(c)
	errors.PanicOnError(os.Remove(path))
}
