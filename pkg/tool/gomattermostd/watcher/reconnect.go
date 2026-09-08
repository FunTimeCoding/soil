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

	select {
	case <-w.done:
		return false
	case <-time.After(constant.ReconnectDelay):
	}

	w.client.RefreshSocket()
	w.client.WebSocket().Listen()

	return true
}
