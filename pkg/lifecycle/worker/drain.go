package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"time"
)

func (w *Worker) Drain(deadline time.Time) {
	d, okay := w.Face.(face.Drainer)

	if !okay {
		return
	}

	select {
	case <-d.Finished():
		return
	case <-time.After(time.Until(deadline)):
	}

	d.Interrupt()

	select {
	case <-d.Finished():
	case <-time.After(constant.InterruptMargin):
	}
}
