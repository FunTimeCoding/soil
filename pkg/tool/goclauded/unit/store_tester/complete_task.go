package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) CompleteTask(name string) string {
	result, e := o.Store.CompleteTask(name)
	assert.FatalOnError(o.t, e)

	return result
}
