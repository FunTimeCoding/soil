package run

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/run/process"
)

func (r *Run) Begin(s ...string) face.Process {
	c := r.build(s...)
	p := process.NewPiped(c)
	errors.PanicOnError(c.Start())

	return p
}
