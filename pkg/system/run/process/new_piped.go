package process

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os/exec"
)

func NewPiped(c *exec.Cmd) *Process {
	stdout, e := c.StdoutPipe()
	errors.PanicOnError(e)
	stderr, f := c.StderrPipe()
	errors.PanicOnError(f)

	return &Process{command: c, stdout: stdout, stderr: stderr}
}
