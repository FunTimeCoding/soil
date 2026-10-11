package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) QueueRows() int64 {
	o.t.Helper()
	var result int64
	assert.FatalOnError(
		o.t,
		o.Store.Database().Raw("SELECT COUNT(*) FROM queue").Scan(
			&result,
		).Error,
	)

	return result
}
