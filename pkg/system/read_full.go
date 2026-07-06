package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func ReadFull(
	r io.Reader,
	b []byte,
) int {
	result, e := io.ReadFull(r, b)
	errors.PanicOnError(e)

	return result
}
