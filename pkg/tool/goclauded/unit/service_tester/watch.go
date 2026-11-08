package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (o *Tester) Watch(sessionIdentifier string) string {
	callsign := o.Check(sessionIdentifier).Callsign
	_, e := o.Label.SetLabel(
		sessionIdentifier,
		callsign,
		callsign,
		constant.CompactionWatchLabel,
		constant.CompactionWatchValue,
	)
	assert.FatalOnError(o.t, e)

	return callsign
}
