package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/console_tester"
	"strings"
	"testing"
	"time"
)

func TestRegisterReturnsCallsign(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	callsign := c.Register("session-1")
	assert.True(t, callsign != "")
}

func TestRegisterSameSessionReturnsSameCallsign(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	first := c.Register("session-1")
	second := c.Register("session-1")
	assert.String(t, first, second)
}

func TestCheckEmptyQueueReturnsEmpty(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-1")
	output := c.Check("session-1")
	assert.String(t, "", output)
}

func TestHookCheckShowsSessionActivity(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-2")
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	output := c.Check("session-2")
	assert.StringContains(t, "Session activity", output)
	assert.StringContains(t, a.Name(), output)
}

func TestHookCheckConsumesEntries(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-2")
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	first := c.Check("session-2")
	assert.True(t, first != "")
	second := c.Check("session-2")
	assert.String(t, "", second)
}

func TestCheckTimeout(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
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
	c := console_tester.New(t, s.Client(t))
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
	c := console_tester.New(t, s.Client(t))
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
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	c := console_tester.New(t, s.Client(t))
	c.Register("session-2")
	c.Check("session-2")
	s.Service.ClearBindings()
	output := c.Check(a.UUID)
	assert.StringContains(t, "Re-announce", output)
}

func TestStatusLineRendersAndStores(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	line := c.StatusLine([]byte(fixture.Read("claude", "status-line.json")))
	assert.String(t, "Fable 5 18%", line)
	record, found, e := s.Service.FindSession(
		"11111111-2222-3333-4444-555555555555",
	)
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.Integer(t, 18, record.ContextPercent)
	assert.Integer(t, 1000000, record.ContextWindow)
	assert.String(t, "Fable 5", record.Model)
}

func TestStatusLineKeepsUnmappedModelName(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		"Fable 5",
		"Opus 4.6",
	)
	assert.String(t, "Opus 4.6 18%", c.StatusLine([]byte(body)))
}

func TestStatusLineRecordsFableFromModelScope(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	c.StatusLine([]byte(fixture.Read("claude", "status-line.json")))
	snapshot, e := s.Store.Store.LatestFableSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, snapshot)
	assert.Integer(t, 34, snapshot.Percent)
	assert.String(t, "", snapshot.Reset)
	assert.FatalNotNil(t, snapshot.ResetAt)
	assert.String(
		t,
		"2023-11-15T00:00:00Z",
		snapshot.ResetAt.UTC().Format(time.RFC3339),
	)
}

func TestStatusLineWithoutModelScopeRecordsNoFable(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		`"display_name": "Fable"`,
		`"display_name": "Opus"`,
	)
	c.StatusLine([]byte(body))
	snapshot, e := s.Store.Store.LatestFableSnapshot()
	assert.FatalOnError(t, e)
	assert.Nil(t, snapshot)
}

func TestStatusLineLeavesMissingResetZero(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		`"resets_at": 1700003600`,
		`"resets_at": 0`,
	)
	c.StatusLine([]byte(body))
	snapshot, e := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, snapshot)
	assert.True(t, snapshot.FiveHourReset.IsZero())
	assert.False(t, snapshot.SevenDayReset.IsZero())
}

func TestStatusLineRecordsWeekWithoutFiveHour(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		`"five_hour": {
      "used_percentage": 31,
      "resets_at": 1700003600
    },`,
		"",
	)
	assert.StringNotContains(t, "five_hour", body)
	c.StatusLine([]byte(body))
	snapshot, e := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, snapshot)
	assert.Integer(t, 1, snapshot.SevenDayPercent)
	assert.Integer(t, 0, snapshot.FiveHourPercent)
	assert.True(t, snapshot.FiveHourReset.IsZero())
}

func TestStatusLineSkipsSnapshotWithoutSevenDay(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		`"seven_day": {
      "used_percentage": 1,
      "resets_at": 1700006400
    },`,
		"",
	)
	assert.StringNotContains(t, "seven_day", body)
	c.StatusLine([]byte(body))
	snapshot, e := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.Nil(t, snapshot)
}

func TestStatusLineRateSnapshotDedupe(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("11111111-2222-3333-4444-555555555555")
	body := []byte(fixture.Read("claude", "status-line.json"))
	c.StatusLine(body)
	first, e := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.NotNil(t, first)
	assert.Integer(t, 31, first.FiveHourPercent)
	assert.Integer(t, 1, first.SevenDayPercent)
	c.StatusLine(body)
	second, f := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, f)
	assert.Integer(t, int(first.Identifier), int(second.Identifier))
}

func TestTurnEndThroughTheHookCommand(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-1")
	assert.True(t, s.Store.GetSession("session-1").LastTurnEndAt == nil)
	c.TurnEnd("session-1")
	assert.True(t, s.Store.GetSession("session-1").LastTurnEndAt != nil)
}

func TestHookCommandOrderingDecidesIdle(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-1")
	c.Check("session-1")
	s.Store.Advance(time.Minute)
	c.TurnEnd("session-1")
	r := s.Store.GetSession("session-1")
	assert.True(t, r.LastTurnEndAt.After(*r.LastPromptAt))
}

func TestHookCommandOrderingDecidesWorking(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Client(t))
	c.Register("session-1")
	c.TurnEnd("session-1")
	s.Store.Advance(time.Minute)
	c.Check("session-1")
	r := s.Store.GetSession("session-1")
	assert.True(t, r.LastPromptAt.After(*r.LastTurnEndAt))
}
