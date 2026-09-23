package errors

import "io"

func PanicClose(c io.Closer) {
	if e := c.Close(); e != nil {
		panic(e)
	}
}
