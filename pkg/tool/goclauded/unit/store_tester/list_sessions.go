package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (o *Tester) ListSessions() []session.Session {
	result, e := o.Store.ListSessions()
	assert.FatalOnError(o.t, e)

	return result
}
