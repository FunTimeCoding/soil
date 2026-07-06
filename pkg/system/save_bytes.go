package system

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
)

func SaveBytes(
	path string,
	b []byte,
) {
	f := Create(path)
	w := bufio.NewWriter(f)
	_, e := w.Write(b)
	errors.PanicOnError(e)
	errors.PanicOnError(w.Flush())
	errors.PanicClose(f)
}
