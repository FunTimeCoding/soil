package console

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/console_tester"
	"testing"
	"time"
)

func TestRegisterReturnsCallsign(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	callsign := c.Register("session-1")
	assert.True(t, callsign != "")
}

func TestRegisterSameSessionReturnsSameCallsign(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	first := c.Register("session-1")
	second := c.Register("session-1")
	assert.String(t, first, second)
}

func TestCheckEmptyQueueReturnsEmpty(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("session-1")
	output := c.Check("session-1")
	assert.String(t, "", output)
}

func TestCheckWithSessionActivity(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("session-2")
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	output := c.Check("session-2")
	assert.StringContains(t, "Session activity", output)
	assert.StringContains(t, a.Name(), output)
}

func TestCheckConsumesEntries(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("session-2")
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	first := c.Check("session-2")
	assert.True(t, first != "")
	second := c.Check("session-2")
	assert.String(t, "", second)
}

func TestCheckTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	callsign := c.Register("session-1")
	s.Announce("session-1", callsign, "working", "")
	c.Check("session-1")
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	output := c.Check("session-1")
	assert.StringContains(
		t,
		"Idle: 1 hour since last turn. Removed from roster.",
		output,
	)
}

func TestCheckDirectMessage(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	receiver := c.Register("session-1")
	sender := c.Register("session-2")
	c.Check("session-1")
	s.Send(sender, receiver, "heads up: deploying")
	output := c.Check("session-1")
	assert.StringContains(t, "Messages", output)
	assert.StringContains(t, "deploying", output)
}

func TestCheckCompletionActivity(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("session-1")
	other := c.Register("session-2")
	s.Announce("session-2", other, "indexing", "")
	c.Check("session-1")
	s.Complete("session-2", other, "indexing", "done")
	output := c.Check("session-1")
	assert.StringContains(t, "Session activity", output)
	assert.StringContains(t, other, output)
	assert.StringContains(t, "completed", output)
}

func TestCheckReannounceAfterClearBindings(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	c := console_tester.New(t, s.Port())
	c.Register("session-2")
	c.Check("session-2")
	s.Service.ClearBindings()
	output := c.Check(a.UUID)
	assert.StringContains(t, "Re-announce", output)
}
