package stream

import "io"

func Drain(r io.ReadCloser) error {
	return r.Close()
}
