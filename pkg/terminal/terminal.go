package terminal

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
)

type Terminal struct {
	ender   face.CommandEnder
	output  io.Writer
	failure io.Writer
	exit    func(int)
}
