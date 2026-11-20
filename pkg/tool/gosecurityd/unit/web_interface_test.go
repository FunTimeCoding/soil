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

func TestFlagsPageListsFlagWithDebianStatus(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.KnownIdentifier,
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
			"identifier":       {vulnerabilityConstant.KnownIdentifier},
			constant.NoteField: {"india"},
		},
	)
	assert.StringContains(t, "Flagged CVE-2024-3094", answer)
	assert.Eventually(
		t,
		func() bool {
			return strings.Contains(
				v.Body(t, constant.FlagsPath),
				"trixie fixed",
			)
		},
	)
	b := v.Body(t, constant.FlagsPath)
	assert.StringContains(t, "CVE-2024-3094", b)
	assert.StringContains(t, "CRITICAL 10.0", b)
	assert.StringContains(t, "india", b)
	assert.StringContains(t, "bookworm not affected", b)
	assert.StringContains(t, `title="xz-utils 5.6.1-1"`, b)
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
		"vulnerability identifier malformed: juliett",
		answer,
	)
	assert.StringContains(t, `value="juliett"`, answer)
}

func TestMitigateRemovesRowFromFlagsPage(t *testing.T) {
	v := base.New(t).Served()
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {vulnerabilityConstant.KnownIdentifier}},
	)
	answer := v.PostBody(
		t,
		constant.MitigatePath,
		url.Values{
			"identifier": {vulnerabilityConstant.KnownIdentifier},
			"mitigation": {"kilo"},
		},
	)
	assert.StringContains(t, "Mitigated CVE-2024-3094", answer)
	assert.StringContains(t, "No flags here.", answer)
}

func TestMitigatingUnflaggedShowsRefusalInline(t *testing.T) {
	s := base.New(t)
	answer := s.Served().PostBody(
		t,
		constant.MitigatePath,
		url.Values{"identifier": {vulnerabilityConstant.KnownIdentifier}},
	)
	assert.StringContains(t, "flag not found: CVE-2024-3094", answer)
	assert.StringContains(t, "No flags here.", answer)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestLookupShowsBothSources(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.KnownIdentifier,
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
			"identifier": {vulnerabilityConstant.KnownIdentifier},
		}.Encode(),
	)
	assert.Eventually(
		t,
		func() bool {
			return strings.Contains(v.Body(t, target), "fixed in 5.6.1-1")
		},
	)
	b := v.Body(t, target)
	assert.StringContains(t, "CRITICAL 10.0", b)
	assert.StringContains(t, "never affected", b)
	assert.StringContains(t, `value="CVE-2024-3094"`, b)
}
