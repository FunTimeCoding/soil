package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestSend(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r1.Callsign, r2.Callsign, "heads up, touching auth")
	messages := s.Store.PendingMessages(r2.Callsign)
	assert.Count(t, 1, messages)
	assert.String(t, "heads up, touching auth", messages[0].Body)
}

func TestSendBroadcast(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r1.Callsign, "", "deploying now")
	messages := s.Store.PendingMessages(r2.Callsign)
	assert.Count(t, 1, messages)
}
