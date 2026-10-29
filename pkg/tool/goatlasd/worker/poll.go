package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service/outcome"
	"time"
)

func (w *Worker) Poll() {
	started := time.Now()
	v, e := w.service.Collect(w.context)

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
