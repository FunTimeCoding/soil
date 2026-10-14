package errors

import (
	"io"
	"net"
	"os"
)

func LogReclose(c io.Closer) {
	if e := c.Close(); !Is(e, os.ErrClosed) && !Is(e, net.ErrClosed) {
		LogOnError(e)
	}
}
