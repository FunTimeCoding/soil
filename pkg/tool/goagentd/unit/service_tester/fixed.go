package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
)

func Fixed(s face.Stream) func(string, io.Writer, ...string) face.Stream {
	return func(
		string,
		io.Writer,
		...string,
	) face.Stream {
		return s
	}
}
