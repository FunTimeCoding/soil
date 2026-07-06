package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (o *Tester) GetMemoryHistory(identifier int64) []store.Version {
	o.t.Helper()
	result, e := o.Store.GetMemoryHistory(identifier)
	assert.FatalOnError(o.t, e)

	return result
}
