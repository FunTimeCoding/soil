package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/run"
	"io"
)

func New() *Agent {
	return &Agent{
		Stream: func(
			directory string,
			stderr io.Writer,
			s ...string,
		) face.Stream {
			r := run.New()
			r.Directory = directory
			r.Writers(nil, stderr)

			return r.Stream(s...)
		},
	}
}
