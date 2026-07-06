package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SetListening(
	name string,
	listening bool,
) {
	assert.FatalOnError(o.t, o.Store.SetListening(name, listening))
}
