package run

import (
	"github.com/funtimecoding/soil/pkg/system/run/process"
	"os"
)

func (r *Run) TryOpen(s ...string) (*process.Process, error) {
	c := r.build(s...)

	if r.stdio {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	if e := c.Start(); e != nil {
		return nil, e
	}

	return process.New(c), nil
}
