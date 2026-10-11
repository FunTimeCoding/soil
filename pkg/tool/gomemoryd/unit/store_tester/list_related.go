package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (o *Tester) ListRelated(identifier int64) []record.Related {
	o.t.Helper()
	result, e := o.Store.ListRelated(identifier)
	assert.FatalOnError(o.t, e)

	return result
}
