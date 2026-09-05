package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (o *Tester) CompletionsBySession(
	sessionIdentifier string,
) []completion.Completion {
	result, e := o.Store.CompletionsBySession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
