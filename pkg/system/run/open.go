package run

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"os"
)

func (r *Run) Open(s ...string) *process.Process {
	c := r.build(s...)

	if r.stdio {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	errors.PanicOnError(c.Start())

	return process.New(c)
}
