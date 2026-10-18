package watcher

import "github.com/funtimecoding/soil/pkg/peer/reacher/edge"

func (w *Watcher) logEdge(x *edge.Edge) {
	w.logger.Plain("websocket %s", x)
}
