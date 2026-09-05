package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) ReleaseByCallsign(callsign string) {
	assert.FatalOnError(o.t, o.Store.ReleaseByCallsign(callsign))
}
