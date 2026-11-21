package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/url"
	"strings"
	"testing"
)

func TestFlagsRowStaysNarrowAndDetailCarriesDebian(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.FixtureIdentifier,
		"CRITICAL",
		10.0,
	)
	s.SecurityTracker.SeedKnown()
	s.Started()
	v := s.Served()
	answer := v.PostBody(
		t,
		constant.FlagPath,
		url.Values{
			"identifier":       {vulnerabilityConstant.FixtureIdentifier},
			constant.NoteField: {"india"},
		},
	)
	assert.StringContains(t, "Flagged CVE-2099-00001", answer)
	b := v.Body(t, constant.FlagsPath)
	assert.StringContains(t, "CVE-2099-00001", b)
	assert.StringContains(t, "CRITICAL 10.0", b)
	assert.StringContains(
		t,
		"<tr><th>Subject</th><th>Severity</th><th>Hits</th><th>Flagged</th></tr>",
		b,
	)
	assert.StringNotContains(t, "india", b)
	fixed := `<span title="tango-utils fixed in 8.6.1-1"><span class="status-dot" style="--dot-color: var(--okay);"></span> <small>trixie</small></span>`
	opened := join.Empty(
		constant.FlagOpenPath,
		"?",
		url.Values{"identifier": {vulnerabilityConstant.FixtureIdentifier}}.Encode(),
	)
	assert.Eventually(
		t,
		func() bool { return strings.Contains(v.Body(t, opened), fixed) },
	)
	d := v.Body(t, opened)
	assert.StringContains(t, "india", d)
	assert.StringContains(t, `colspan="4"`, d)
	assert.StringContains(
		t,
		"<details><summary>Versions</summary><ul><li>tango-utils",
		d,
	)
	assert.StringContains(
		t,
		`<span title="tango-utils not affected"><span class="status-dot" style="--dot-color: var(--quiet);"></span> <small>bookworm</small></span>`,
		d,
	)
}

func TestFlagFormShowsRefusalInline(t *testing.T) {
	v := base.New(t).Served()
	answer := v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {"juliett"}},
	)
	assert.StringContains(
		t,
		"subject malformed, expect a CVE identifier, upgrade:&lt;release&gt;:&lt;source&gt;, release:&lt;host&gt;, reboot:&lt;host&gt; or kernel:&lt;host&gt;: juliett",
		answer,
	)
	assert.StringContains(t, `value="juliett"`, answer)
}

func TestMitigateRemovesRowFromFlagsPage(t *testing.T) {
	v := base.New(t).Served()
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {vulnerabilityConstant.FixtureIdentifier}},
	)
	answer := v.PostBody(
		t,
		constant.MitigatePath,
		url.Values{
			"identifier": {vulnerabilityConstant.FixtureIdentifier},
			"mitigation": {"kilo"},
		},
	)
	assert.StringContains(t, "Mitigated CVE-2099-00001", answer)
	assert.StringContains(t, "No flags here.", answer)
}

func TestMitigatingUnflaggedShowsRefusalInline(t *testing.T) {
	s := base.New(t)
	answer := s.Served().PostBody(
		t,
		constant.MitigatePath,
		url.Values{"identifier": {vulnerabilityConstant.FixtureIdentifier}},
	)
	assert.StringContains(t, "flag not found: CVE-2099-00001", answer)
	assert.StringContains(t, "No flags here.", answer)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestLookupShowsBothSources(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.FixtureIdentifier,
		"CRITICAL",
		10.0,
	)
	s.SecurityTracker.SeedKnown()
	s.Started()
	v := s.Served()
	target := join.Empty(
		constant.LookupPath,
		"?",
		url.Values{
			"identifier": {vulnerabilityConstant.FixtureIdentifier},
		}.Encode(),
	)
	assert.Eventually(
		t,
		func() bool {
			return strings.Contains(v.Body(t, target), "fixed in 8.6.1-1")
		},
	)
	b := v.Body(t, target)
	assert.StringContains(t, "CRITICAL 10.0", b)
	assert.StringContains(t, "never affected", b)
	assert.StringContains(t, `value="CVE-2099-00001"`, b)
}
