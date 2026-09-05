package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
)

func (o *Tester) EditSession(
	identifier string,
	a *edit_session.Session,
) {
	assert.FatalOnError(o.t, o.Store.EditSession(identifier, a))
}
