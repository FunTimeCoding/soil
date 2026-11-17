package base

import (
	"github.com/funtimecoding/soil/pkg/face"
	"io"
	"testing"
)

func New(
	t *testing.T,
	stream func(
		string,
		io.Writer,
		...string,
	) face.Stream,
) *Server {
	t.Helper()

	return &Server{stream: stream, t: t}
}
