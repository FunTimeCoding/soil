package run

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run/stream"
)

func (r *Run) Stream(s ...string) *stream.Stream {
	c := r.build(s...)
	reader, e := c.StdoutPipe()
	errors.PanicOnError(e)
	errors.PanicOnError(c.Start())

	return stream.New(reader, c)
}
