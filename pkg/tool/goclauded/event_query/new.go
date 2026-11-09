package event_query

import "github.com/funtimecoding/soil/pkg/time/window"

func New() *Query {
	return &Query{window: window.Whole()}
}
