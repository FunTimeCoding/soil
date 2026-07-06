package system

import (
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func GnuZipReader(r io.Reader) *gzip.Reader {
	result, e := gzip.NewReader(r)
	errors.PanicOnError(e)

	return result
}
