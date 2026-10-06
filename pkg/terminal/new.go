package terminal

import (
	"github.com/funtimecoding/soil/pkg/face"
	"os"
)

func New(e face.CommandEnder) *Terminal {
	return NewWith(e, os.Stdout, os.Stderr, os.Exit)
}
