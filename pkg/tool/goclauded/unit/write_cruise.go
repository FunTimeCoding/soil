package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"testing"
	"time"
)

func writeCruise(
	t *testing.T,
	s *store.Store,
	sessionIdentifier string,
	callsign string,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.SetCruise(sessionIdentifier, constant.CruisePaced, time.Minute),
	)
	assert.FatalOnError(t, s.MarkPromoted(sessionIdentifier))
	held := queue.New(sessionIdentifier, callsign, constant.QueuePulse, "held")
	held.Delivery = constant.DeliveryCruise
	assert.FatalOnError(t, s.PushEntry(held))
	released, e := s.ReleaseHeldCruise(sessionIdentifier, 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, int(released))
}
