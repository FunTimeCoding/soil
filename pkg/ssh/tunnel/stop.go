package tunnel

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"syscall"
)

func (t *Tunnel) Stop() {
	if t.process == nil {
		panic("not running")
	}

	e := t.process.Signal(syscall.SIGTERM)

	if e != nil && e.Error() != "os: process already finished" {
		errors.PanicOnError(e)
	}

	<-t.stopped
}
