package server_tester

import (
	"errors"
	soilErrors "github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func release(c io.Closer) {
	if e := c.Close(); !errors.Is(e, io.EOF) {
		soilErrors.LogOnError(e)
	}
}
