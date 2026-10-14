package worker

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"

func (w *Worker) Close() {
	w.logger.Plain(constant.StopMessage)
	w.cancel()
}
