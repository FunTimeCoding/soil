package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) ForgetMemory(
	identifier int64,
	source string,
) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Store.ForgetMemory(identifier, source))
}
