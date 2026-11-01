package runner

import (
	"errors"
	"os"
)

func (r *Runner) kill() {
	if r.current == nil {
		return
	}

	if e := r.current.Kill(); e != nil && !errors.Is(e, os.ErrProcessDone) {
		r.reporter.CaptureException(e)
	}
}
