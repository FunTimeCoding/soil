package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (o *Tester) SessionCompletions(
	sessionIdentifier string,
) []completion.Completion {
	result, e := o.Store.SessionCompletions(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
