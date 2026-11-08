package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (o *Tester) Compact(
	sessionIdentifier string,
	phase constant.CompactionPhase,
	trigger string,
) {
	assert.FatalOnError(
		o.t,
		o.Roster.RecordCompaction(sessionIdentifier, phase, trigger),
	)
}
