package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) AddSession(
	identifier string,
	name string,
	callsign any,
	started string,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.Store.Database().Exec(
			`INSERT INTO session (identifier, name, callsign, started_at)
			VALUES (?, ?, ?, ?)`,
			identifier,
			name,
			callsign,
			started,
		).Error,
	)
}
