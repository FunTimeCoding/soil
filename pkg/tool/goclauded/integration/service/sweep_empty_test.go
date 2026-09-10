package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
	"time"
)

func TestSweepRemovesClosedEmptySessionAtOnce(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	assert.FatalOnError(t, s.Service.CloseSession("husk", "prompt_input_exit"))
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("husk") == nil)
}

func TestSweepKeepsUnmarkedEmptySessionUntilCutoff(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("husk") != nil)
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("husk") == nil)
}

func TestSweepKeepsClosedSessionWithContent(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "spoken", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.FatalOnError(t, s.Service.CloseSession("spoken", "logout"))
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("spoken") != nil)
}

func TestSweepKeepsMetadataGhost(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("ghost")
	s.Store.Store.UpdateFields(
		"ghost",
		map[string]any{"alias": "named-by-hand"},
	)
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("ghost") != nil)
}

func TestSweepKeepsDescribedSession(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("described")
	s.Store.Store.UpdateFields(
		"described",
		map[string]any{"description": "written by hand"},
	)
	s.Store.Advance(25 * time.Hour)
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("described") != nil)
}

func TestSweptSessionReregistersOnResume(t *testing.T) {
	s := service_tester.New(t)
	first := s.Store.EnsureSession("husk")
	assert.FatalOnError(t, s.Service.CloseSession("husk", "logout"))
	s.Service.RunTimeoutSweep()
	assert.True(t, s.Store.GetSession("husk") == nil)
	second := s.Store.EnsureSession("husk")
	r := s.Store.GetSession("husk")
	assert.True(t, r != nil)
	assert.True(t, r.ClosedAt == nil)
	assert.Integer(t, 0, r.TurnCount)
	assert.True(t, first.Callsign != "")
	assert.True(t, second.Callsign != "")
}
