package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) StoredValue(identifier string) string {
	o.t.Helper()
	var result string
	assert.FatalOnError(
		o.t,
		o.Store.Database().Raw(
			"SELECT CAST(last_seen AS TEXT) FROM session WHERE identifier = ?",
			identifier,
		).Scan(&result).Error,
	)

	return result
}
