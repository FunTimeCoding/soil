package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (o *Tester) GetSession(sessionIdentifier string) *session.Session {
	result, _, e := o.Store.FindSession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
