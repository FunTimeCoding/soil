package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SessionSummary(sessionIdentifier string) string {
	result, e := o.Store.SessionSummary(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
