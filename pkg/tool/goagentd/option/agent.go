package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
)

type Agent struct {
	Address       string
	ServiceTokens []string
	Workspace     string
	Stream        func(string, io.Writer, ...string) face.Stream
}
