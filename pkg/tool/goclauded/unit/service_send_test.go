package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestSend(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Check("session-2")
	s.Send(r1.Callsign, r2.Callsign, "heads up, touching auth")
	entries := fixture.EntriesByKind(
		s.Check("session-2").Entries,
		constant.QueueMessage,
	)
	assert.Count(t, 1, entries)
	messages := s.ReadMessages(*entries[0].MessageIdentifier)
	assert.Count(t, 1, messages)
	assert.String(t, "heads up, touching auth", messages[0].Body)
}

func TestSendBroadcast(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	s.Check("session-2")
	s.Send(r1.Callsign, "", "deploying now")
	entries := fixture.EntriesByKind(
		s.Check("session-2").Entries,
		constant.QueueMessage,
	)
	assert.Count(t, 1, entries)
	messages := s.ReadMessages(*entries[0].MessageIdentifier)
	assert.String(t, "deploying now", messages[0].Body)
	assert.String(t, "", messages[0].ToName)
}

func TestSendReachesEveryNamedRecipient(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Check("session-2")
	r3 := s.Check("session-3")
	_, e := s.Service.Send(
		r1.Callsign,
		[]string{r2.Callsign, r3.Callsign},
		"alfa",
		constant.DeliveryQueued,
	)
	assert.FatalOnError(t, e)

	for _, session := range []string{"session-2", "session-3"} {
		assert.Count(
			t,
			1,
			fixture.EntriesByKind(
				s.Check(session).Entries,
				constant.QueueMessage,
			),
		)
	}
}

func TestAnUnknownRecipientRefusesTheWholeSend(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Check("session-2")
	_, e := s.Service.Send(
		r1.Callsign,
		[]string{r2.Callsign, "Nobody"},
		"bravo",
		constant.DeliveryQueued,
	)
	assert.True(t, not_found.Is(e))
	assert.Count(
		t,
		0,
		fixture.EntriesByKind(
			s.Check("session-2").Entries,
			constant.QueueMessage,
		),
	)
}

func TestABroadcastSkipsItsSenderAndQuietSessions(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	s.Check("session-2")
	r3 := s.Check("session-3")
	_, e := s.Service.SetLabel(
		"session-3",
		r3.Callsign,
		r3.Callsign,
		constant.BroadcastLabel,
		constant.BroadcastOff,
	)
	assert.FatalOnError(t, e)
	s.Send(r1.Callsign, "", "charlie")
	count := func(session string) int {
		return len(
			fixture.EntriesByKind(
				s.Check(session).Entries,
				constant.QueueMessage,
			),
		)
	}
	assert.Integer(t, 0, count("session-1"))
	assert.Integer(t, 1, count("session-2"))
	assert.Integer(t, 0, count("session-3"))
}

func TestANamedQuietSessionIsStillReached(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Check("session-2")
	_, e := s.Service.SetLabel(
		"session-2",
		r2.Callsign,
		r2.Callsign,
		constant.BroadcastLabel,
		constant.BroadcastOff,
	)
	assert.FatalOnError(t, e)
	s.Send(r1.Callsign, r2.Callsign, "delta")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(
			s.Check("session-2").Entries,
			constant.QueueMessage,
		),
	)
}

func TestImmediateWakesEveryNamedRecipient(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Check("session-2")
	r3 := s.Check("session-3")
	held, e := s.Service.Send(
		r1.Callsign,
		[]string{r2.Callsign, r3.Callsign},
		"echo",
		constant.DeliveryImmediate,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, held)
	two, f := s.Service.DrainImmediateQueue("session-2", r2.Callsign)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, two)
	three, g := s.Service.DrainImmediateQueue("session-3", r3.Callsign)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, three)
}

func TestReadMessagesSkipsUnknown(t *testing.T) {
	s := service_tester.New(t)
	assert.Count(t, 0, s.ReadMessages(404))
}
