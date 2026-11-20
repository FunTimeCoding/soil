package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestFlagToolsRoundTrip(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.KnownIdentifier,
		"CRITICAL",
		10.0,
	)
	c := s.Tools()
	f := c.MustCallToolNotation(
		constant.FlagVulnerability,
		map[string]any{
			"identifier":       vulnerabilityConstant.KnownIdentifier,
			constant.NoteField: "golf",
		},
	)
	assert.MapValue(t, "CVE-2024-3094", f, "identifier")
	assert.MapValue(t, "CRITICAL", f, "severity")
	assert.MapValue(t, false, f, "mitigated")
	open := c.MustCallToolNotation(constant.ListFlags, map[string]any{})
	assert.Count(t, 1, open["flags"])
	m := c.MustCallToolNotation(
		constant.MitigateFlag,
		map[string]any{
			"identifier":           vulnerabilityConstant.KnownIdentifier,
			"mitigation": "hotel",
		},
	)
	assert.MapValue(t, true, m, "mitigated")
	assert.MapValue(t, "hotel", m, "mitigation")
	after := c.MustCallToolNotation(constant.ListFlags, map[string]any{})
	assert.Count(t, 0, after["flags"])
	all := c.MustCallToolNotation(
		constant.ListFlags,
		map[string]any{"include_mitigated": true},
	)
	assert.Count(t, 1, all["flags"])
}

func TestFlaggingTwiceIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	c.MustCallToolNotation(
		constant.FlagVulnerability,
		map[string]any{"identifier": vulnerabilityConstant.KnownIdentifier},
	)
	assert.String(
		t,
		"flag already exists: CVE-2024-3094",
		c.MustCallToolError(
			constant.FlagVulnerability,
			map[string]any{"identifier": vulnerabilityConstant.KnownIdentifier},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
