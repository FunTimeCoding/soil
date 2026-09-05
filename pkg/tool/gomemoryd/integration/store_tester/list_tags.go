package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (o *Tester) ListTags() []store.TagCount {
	o.t.Helper()
	result, e := o.Store.ListTags()
	assert.FatalOnError(o.t, e)

	return result
}
