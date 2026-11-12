package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) PendingQueue(
	sessionIdentifier string,
	body string,
) int64 {
	o.t.Helper()
	var result int64
	assert.FatalOnError(
		o.t,
		o.Store.Database().Raw(
			`SELECT COUNT(*) FROM queue
			WHERE consumed_at IS NULL
			AND COALESCE(session_identifier, '') = ?
			AND body = ?`,
			sessionIdentifier,
			body,
		).Scan(&result).Error,
	)

	return result
}
