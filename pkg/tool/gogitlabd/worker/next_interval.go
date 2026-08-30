package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"time"
)

func (w *Worker) nextInterval() time.Duration {
	if w.Active() {
		return constant.ActivePollInterval
	}

	return w.interval
}
