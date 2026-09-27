package console

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/console_tester"
	"strings"
	"testing"
	"time"
)

func TestStatusLineRendersAndStores(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
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
	c := console_tester.New(t, s.Port)
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
	c := console_tester.New(t, s.Port)
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
	c := console_tester.New(t, s.Port)
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

func TestStatusLineStoresResetsAsUniversal(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
	c.Register("11111111-2222-3333-4444-555555555555")
	c.StatusLine([]byte(fixture.Read("claude", "status-line.json")))
	snapshot, e := s.Store.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, snapshot)
	assert.String(t, "+00:00", snapshot.FiveHourReset.Format("-07:00"))
	assert.String(t, "+00:00", snapshot.SevenDayReset.Format("-07:00"))
}

func TestStatusLineLeavesMissingResetZero(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
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

func TestStatusLineSkipsSnapshotWhenWindowMissing(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
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
	assert.Nil(t, snapshot)
}

func TestStatusLineRateSnapshotDedupe(t *testing.T) {
	s := base.New(t)
	c := console_tester.New(t, s.Port)
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
