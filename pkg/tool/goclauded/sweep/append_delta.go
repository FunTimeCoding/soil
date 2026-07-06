package sweep

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func appendDelta(
	source string,
	destination string,
	offset int64,
) {
	s, e := os.Open(source)
	errors.PanicOnError(e)
	defer errors.PanicClose(s)
	_, e = s.Seek(offset, io.SeekStart)
	errors.PanicOnError(e)
	d, f := os.OpenFile(destination, os.O_APPEND|os.O_WRONLY, 0644)
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
