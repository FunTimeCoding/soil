package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (o *Tester) Events(q *event_query.Query) []event.Event {
	result, e := o.Store.Events(q)
	assert.FatalOnError(o.t, e)

	return result
}
