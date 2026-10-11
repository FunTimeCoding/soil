package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) AddQueueRow(
	callsign any,
	body string,
	created string,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.Database().Exec(
			`INSERT INTO queue (callsign, kind, body, consumed, created_at)
			VALUES (?, 'message', ?, 0, ?)`,
			callsign,
			body,
			created,
		).Error,
	)
}
