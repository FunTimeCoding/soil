package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (o *Tester) ListTags() []record.TagCount {
	o.t.Helper()
	result, e := o.Store.ListTags()
	assert.FatalOnError(o.t, e)

	return result
}
