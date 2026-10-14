package lifecycle

import (
	"github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"time"
)

func (l *Lifecycle) Close() {
	deadline := time.Now().Add(l.grace - constant.InterruptMargin)

	for _, w := range l.drainer {
		w.Close()
	}

	for i := len(l.component) - 1; i >= 0; i-- {
		l.component[i].Close()
	}

	for _, w := range l.drainer {
		w.Drain(deadline)
	}
}
