package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestClientReadsVulnerability(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.KnownIdentifier,
		"CRITICAL",
		10.0,
	)
	v, e := s.SecurityClient().GetVulnerability(
		t.Context(),
		vulnerabilityConstant.KnownIdentifier,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "CVE-2024-3094", v.Identifier)
	assert.String(t, "CRITICAL", v.Severity)
	assert.Float(t, 10.0, v.Score)
}

func TestClientReadsDebianStatus(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	release := "trixie"
	var r *client.DebianStatus
	assert.Eventually(
		t,
		func() bool {
			var e error
			r, e = c.GetDebianStatus(
				t.Context(),
				vulnerabilityConstant.KnownIdentifier,
				&client.GetDebianStatusParams{Release: &release},
			)

			return e == nil
		},
	)
	assert.Count(t, 1, r.Statuses)
	assert.String(t, "trixie", r.Statuses[0].Release)
	assert.String(t, "5.6.1-1", *r.Statuses[0].FixedVersion)
}

func TestClientRoundTripsFlag(t *testing.T) {
	s := base.New(t)
	s.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.KnownIdentifier,
		"CRITICAL",
		10.0,
	)
	c := s.SecurityClient()
	x := t.Context()
	note := "lima"
	f, e := c.FlagVulnerability(
		x,
		client.FlagRequest{
			Identifier: vulnerabilityConstant.KnownIdentifier,
			Note:       &note,
		},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "CVE-2024-3094", f.Identifier)
	assert.String(t, "CRITICAL", *f.Severity)
	assert.String(t, "lima", *f.Note)
	open, g := c.ListFlags(x, &client.ListFlagsParams{})
	assert.FatalOnError(t, g)
	assert.Count(t, 1, open)
	how := "mike"
	m, h := c.MitigateFlag(
		x,
		vulnerabilityConstant.KnownIdentifier,
		client.MitigationRequest{Mitigation: &how},
	)
	assert.FatalOnError(t, h)
	assert.True(t, m.Mitigated)
	assert.String(t, "mike", *m.Mitigation)
	after, i := c.ListFlags(x, &client.ListFlagsParams{})
	assert.FatalOnError(t, i)
	assert.Count(t, 0, after)
	all := true
	every, j := c.ListFlags(x, &client.ListFlagsParams{IncludeMitigated: &all})
	assert.FatalOnError(t, j)
	assert.Count(t, 1, every)
}

func TestClientReachesDaemon(t *testing.T) {
	assert.FatalOnError(t, base.New(t).SecurityClient().Reachable(t.Context()))
}
