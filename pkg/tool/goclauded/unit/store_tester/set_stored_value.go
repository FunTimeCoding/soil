package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SetStoredValue(
	identifier string,
	value string,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.Database().Exec(
			"UPDATE session SET last_seen = ? WHERE identifier = ?",
			value,
			identifier,
		).Error,
	)
}
