package face

import "io"

type Stream interface {
	Reader() io.ReadCloser
	Kill() error
	Wait() error
}
