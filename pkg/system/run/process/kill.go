package process

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

func (p *Process) Kill() error {
	if e := p.command.Process.Signal(os.Kill); e != nil {
		return e
	}

	e := p.command.Wait()

	if exit, okay := errors.AsType[*exec.ExitError](e); okay &&
		exit.Sys().(syscall.WaitStatus).Signaled() {
		return nil
	}

	return e
}
