package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (o *Tester) ListMemories(
	memoryType string,
	tag string,
	activeOnly bool,
) []store.MemorySummary {
	o.t.Helper()
	result, e := o.Store.ListMemories(memoryType, tag, activeOnly)
	assert.FatalOnError(o.t, e)

	return result
}
