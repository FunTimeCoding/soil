package sweep

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func copyFile(
	source string,
	destination string,
) {
	s, e := os.Open(source)
	errors.PanicOnError(e)
	defer errors.PanicClose(s)
	d, f := os.Create(destination)
	errors.PanicOnError(f)
	defer errors.PanicClose(d)
	_, g := io.Copy(d, s)
	errors.PanicOnError(g)
	i, h := os.Stat(source)
	errors.PanicOnError(h)
	errors.PanicOnError(
		os.Chtimes(destination, i.ModTime(), i.ModTime()),
	)
}
