package event_query

import "github.com/funtimecoding/soil/pkg/time/window"

func (q *Query) SetWindow(w *window.Window) *Query {
	q.window = w

	return q
}
