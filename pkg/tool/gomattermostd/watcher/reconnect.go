package watcher

import (
	"fmt"
	reacherConstant "github.com/funtimecoding/soil/pkg/reacher/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"time"
)

func (w *Watcher) reconnect(reason string) bool {
	if edge := w.reacher.Fail(w.host, reason); edge != nil {
		w.reporter.CaptureWithContext(
			fmt.Errorf("websocket %s, reconnecting", reason),
			reacherConstant.ContextKey,
			edge.Context(),
		)
	}

	for {
		select {
		case <-w.done:
			return false
		case <-time.After(constant.ReconnectDelay):
		}

		if e := w.client.RefreshSocket(); e != nil {
			w.reacher.Observe(w.host, e)

			continue
		}

		if edge := w.reacher.Succeed(w.host); edge != nil {
			w.logger.Plain("%s", edge)
		}

		w.client.WebSocket().Listen()

		return true
	}
}
