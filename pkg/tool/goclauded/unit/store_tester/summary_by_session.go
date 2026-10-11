package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SummaryBySession(sessionIdentifier string) string {
	result, e := o.Store.SummaryBySession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
