package run

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (r *Run) Execute(s ...string) {
	c := r.build(s...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	e := c.Run()
	r.Error = e

	if r.Panic {
		errors.PanicOnError(e)
	}
}
