package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (o *Tester) UpdateMemory(
	identifier int64,
	option *save_option.Option,
) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Store.UpdateMemory(identifier, option))
}
