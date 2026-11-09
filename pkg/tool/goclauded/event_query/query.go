package event_query

import "github.com/funtimecoding/soil/pkg/time/window"

type Query struct {
	window *window.Window
	kinds  []string
	limit  int
	offset int
}
