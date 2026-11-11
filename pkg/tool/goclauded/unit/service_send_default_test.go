package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestSendWithoutDeliveryWakesNamedRecipient(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	_, e := s.Message.Send(r2.Callsign, []string{r1.Callsign}, "india", "")
	assert.FatalOnError(t, e)
	drained, f := s.Message.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, drained)
}

func TestBroadcastWithoutDeliveryStaysQueued(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	_, e := s.Message.Send(r2.Callsign, nil, "juliett", "")
	assert.FatalOnError(t, e)
	drained, f := s.Message.DrainImmediateQueue("session-1", r1.Callsign)
	assert.FatalOnError(t, f)
	assert.Count(t, 0, drained)
	remaining := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.KindEntries(remaining.Entries, constant.QueueMessage),
	)
}
