package console

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/console_tester"
	"strings"
	"testing"
)

func TestStatusLineRendersAndStores(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("11111111-2222-3333-4444-555555555555")
	line := c.StatusLine([]byte(fixture.Read("claude", "status-line.json")))
	assert.String(t, "Fable 18%", line)
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
	defer s.Close()
	c := console_tester.New(t, s.Port())
	c.Register("11111111-2222-3333-4444-555555555555")
	body := strings.ReplaceAll(
		fixture.Read("claude", "status-line.json"),
		"Fable 5",
		"Opus 4.6",
	)
	assert.String(t, "Opus 4.6 18%", c.StatusLine([]byte(body)))
}

func TestStatusLineRateSnapshotDedupe(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := console_tester.New(t, s.Port())
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
