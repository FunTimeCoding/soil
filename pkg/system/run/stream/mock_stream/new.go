package mock_stream

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
)

func New(
	lines []string,
	open bool,
	ending error,
) *Stream {
	r, w := io.Pipe()
	go func() {
		for _, l := range lines {
			if _, e := io.WriteString(w, join.Empty(l, "\n")); e != nil {
				return
			}
		}

		if !open {
			errors.PanicOnError(w.Close())
		}
	}()

	return &Stream{reader: r, writer: w, ending: ending}
}
