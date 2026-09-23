package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) SessionKeyFor(body string) string {
	o.t.Helper()
	var result string
	assert.FatalOnError(
		o.t,
		o.Store.Database().Raw(
			`SELECT COALESCE(session_identifier, '') FROM queue
			WHERE body = ?`,
			body,
		).Scan(&result).Error,
	)

	return result
}
