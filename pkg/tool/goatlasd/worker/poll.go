package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/outcome"
	"time"
)

func (w *Worker) Poll() {
	started := time.Now()
	v, e := w.collect(w.context)

	if e != nil {
		w.record(
			outcome.Failed(constant.GazetteerSource, time.Since(started), e),
		)

		return
	}

	for _, o := range v {
		w.record(o)
	}
}
