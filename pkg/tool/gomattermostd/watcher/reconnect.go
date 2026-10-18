package watcher

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"time"
)

func (w *Watcher) reconnect(reason string) bool {
	if x := w.reacher.Fail(w.host, reason); x != nil {
		w.logEdge(x)
	}

	for {
		select {
		case <-w.done:
			return false
		case <-time.After(constant.ReconnectDelay):
		}

		e := w.client.RefreshSocket()
		w.reacher.Handle(
			w.host,
			e,
			w.logEdge,
			func(f error) { w.reporter.CaptureException(f) },
		)

		if e != nil {
			continue
		}

		w.client.WebSocket().Listen()

		return true
	}
}
