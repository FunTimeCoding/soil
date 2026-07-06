package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (o *Tester) CreateMemory(p *save_option.Option) int64 {
	o.t.Helper()
	result, e := o.Store.CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return result
}
