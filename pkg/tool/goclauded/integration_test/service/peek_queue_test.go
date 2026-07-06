//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestPeekQueueDoesNotConsume(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r.Callsign, "hello")
	peeked, e := s.Service.PeekQueue(r.Callsign)
	assert.FatalOnError(t, e)
	messages := entriesByKind(peeked, constant.QueueMessage)
	assert.Count(t, 1, messages)
	drained := s.Check("session-1")
	drainedMessages := entriesByKind(
		drained.Entries,
		constant.QueueMessage,
	)
	assert.Count(t, 1, drainedMessages)
}

func TestPeekQueueEmpty(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	peeked, e := s.Service.PeekQueue(r.Callsign)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, peeked)
}
