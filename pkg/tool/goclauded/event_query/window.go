package event_query

import "github.com/funtimecoding/soil/pkg/time/window"

func (q *Query) Window() *window.Window { return q.window }
