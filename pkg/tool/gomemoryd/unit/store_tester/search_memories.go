package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (o *Tester) SearchMemories(
	query string,
	limit int,
	memoryType string,
	tag string,
	scope string,
) []record.SearchResult {
	o.t.Helper()
	result, e := o.Store.SearchMemories(query, limit, memoryType, tag, scope)
	assert.FatalOnError(o.t, e)

	return result
}
