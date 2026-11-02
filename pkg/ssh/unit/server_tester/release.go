package server_tester

import (
	"errors"
	soil "github.com/funtimecoding/soil/pkg/errors"
	"io"
)

func release(c io.Closer) {
	if e := c.Close(); !errors.Is(e, io.EOF) {
		soil.LogOnError(e)
	}
}
