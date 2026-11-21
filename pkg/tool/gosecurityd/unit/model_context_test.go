package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
	"time"
)

func TestGetVulnerabilityAnswersSlimRecord(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.FixtureIdentifier,
		"CRITICAL",
		10.0,
	)
	r := s.Tools().MustCallToolNotation(
		constant.GetVulnerability,
		map[string]any{"identifier": vulnerabilityConstant.FixtureIdentifier},
	)
	assert.MapValue(t, "CVE-2099-00001", r, "identifier")
	assert.MapValue(t, "CRITICAL", r, "severity")
	assert.MapValue(t, 10.0, r, constant.ScoreField)
	assert.MapNotHasKey(t, r, "Raw")
}

func TestGetDebianStatusAnswersPerRelease(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 52, 16, 0, time.UTC),
	)
	s.Started()
	c := s.Tools()
	var r map[string]any
	assert.Eventually(
		t,
		func() bool {
			var e error
			r, e = c.CallToolNotation(
				constant.GetDebianStatus,
				map[string]any{
					"identifier": vulnerabilityConstant.FixtureIdentifier,
				},
			)

			return e == nil
		},
	)
	assert.MapValue(t, "2026-11-19T08:52:16Z", r, "tracker_modified")
	statuses := r["statuses"].([]any)
	assert.Count(t, 2, statuses)
	bookworm := statuses[0].(map[string]any)
	assert.MapValue(t, "bookworm", bookworm, constant.ReleaseParameter)
	assert.MapValue(t, true, bookworm, "never_affected")
	assert.MapNotHasKey(t, bookworm, "fixed_version")
	trixie := statuses[1].(map[string]any)
	assert.MapValue(t, "8.6.1-1", trixie, "fixed_version")
	assert.MapValue(t, false, trixie, "never_affected")
}

func TestUnknownModificationTimeIsLeftOut(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.Tools()
	var r map[string]any
	assert.Eventually(
		t,
		func() bool {
			var e error
			r, e = c.CallToolNotation(
				constant.GetDebianStatus,
				map[string]any{
					"identifier": vulnerabilityConstant.FixtureIdentifier,
				},
			)

			return e == nil
		},
	)
	assert.MapNotHasKey(t, r, "tracker_modified")
}

func TestMalformedIdentifierIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	assert.String(
		t,
		"vulnerability identifier malformed: echo",
		s.Tools().MustCallToolError(
			constant.GetVulnerability,
			map[string]any{"identifier": "echo"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
	assert.Integer(t, 0, s.NationalDatabase.Calls())
}

func TestTrackerNotLoadedIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	assert.StringContains(
		t,
		"not loaded yet",
		s.Tools().MustCallToolError(
			constant.GetDebianStatus,
			map[string]any{
				"identifier": vulnerabilityConstant.FixtureIdentifier,
			},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
