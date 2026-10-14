package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"time"
)

func (o *Tester) EndTurn(sessionIdentifier string) {
	o.Store.Advance(time.Minute)
	assert.FatalOnError(o.t, o.Service.StampTurnEnd(sessionIdentifier))
}
