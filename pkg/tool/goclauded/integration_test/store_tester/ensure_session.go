package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/ensure_result"
)

func (o *Tester) EnsureSession(sessionIdentifier string) *ensure_result.Result {
	result, e := o.Store.EnsureSession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
