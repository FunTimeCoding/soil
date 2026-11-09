package store_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) ReleaseCallsign(callsign string) {
	assert.FatalOnError(o.t, o.Store.ReleaseCallsign(callsign))
}
