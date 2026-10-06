package terminal

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
)

func NewWith(
	e face.CommandEnder,
	output io.Writer,
	failure io.Writer,
	exit func(int),
) *Terminal {
	return &Terminal{ender: e, output: output, failure: failure, exit: exit}
}
