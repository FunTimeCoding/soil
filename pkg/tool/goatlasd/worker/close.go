package worker

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"

func (w *Worker) Close() {
	w.cancel()
	w.started.Do(func() { close(w.done) })
	<-w.done
	w.logger.Plain(constant.StopMessage)
}
