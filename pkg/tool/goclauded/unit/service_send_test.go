package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
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

func TestReadMessagesSkipsUnknown(t *testing.T) {
	s := service_tester.New(t)
	assert.Count(t, 0, s.ReadMessages(404))
}
