package system

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
)

func SaveFile(
	path string,
	text string,
) {
	f := Create(path)
	w := bufio.NewWriter(f)
	_, e := w.WriteString(text)
	errors.PanicOnError(e)
	errors.PanicOnError(w.Flush())
	errors.PanicClose(f)
}
