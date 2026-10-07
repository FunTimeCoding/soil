package watcher

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"time"
)

func (w *Watcher) reconnect(reason string) bool {
	w.reporter.CaptureException(
		fmt.Errorf("websocket %s, reconnecting", reason),
	)

	for {
		select {
		case <-w.done:
			return false
		case <-time.After(constant.ReconnectDelay):
		}

		e := w.client.RefreshSocket()

		if e == nil {
			w.client.WebSocket().Listen()

			return true
		}

		w.logger.Plain("websocket refresh failed: %v", e)
	}
}
