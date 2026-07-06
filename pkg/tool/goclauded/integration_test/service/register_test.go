//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestRegister(t *testing.T) {
	s := service_tester.New(t)
	result := s.Register("session-1")
	assert.True(t, result.New)
	assert.True(t, result.Callsign != "")
}

func TestRegisterExisting(t *testing.T) {
	s := service_tester.New(t)
	first := s.Register("session-1")
	second := s.Register("session-1")
	assert.True(t, first.New)
	assert.False(t, second.New)
	assert.String(t, first.Callsign, second.Callsign)
}

func TestRegisterLogsEvent(t *testing.T) {
	s := service_tester.New(t)
	s.Register("session-1")
	events := s.Store.Events(event_query.New().Kind(constant.Register).SetLimit(10))
	assert.Count(t, 1, events)
}

func TestRegisterDifferentNames(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Register("session-1")
	r2 := s.Register("session-2")
	assert.True(t, r1.Callsign != r2.Callsign)
}

func TestRegisterIncrementsTurnCount(t *testing.T) {
	s := service_tester.New(t)
	s.Register("session-1")
	s.Register("session-1")
	s.Register("session-1")
	assert.True(t, s.Store.GetSession("session-1").TurnCount >= 2)
}
