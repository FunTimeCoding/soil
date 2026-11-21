package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestMockAnswersVulnerabilityLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	m.SeedVulnerability(
		&client.Vulnerability{
			Identifier: vulnerabilityConstant.FixtureIdentifier,
			Severity:   "CRITICAL",
		},
	)
	v, e := m.GetVulnerability(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "CRITICAL", v.Severity)
	_, f := m.GetVulnerability(t.Context(), "CVE-2099-99999")
	assert.True(t, not_found.Is(f))
	_, g := m.GetVulnerability(t.Context(), "oscar")
	assert.True(t, validation.Is(g))
}

func TestMockNarrowsDebianStatusByRelease(t *testing.T) {
	m := mock_client.New(t)
	m.SeedDebianStatus(
		&client.DebianStatus{
			Identifier: vulnerabilityConstant.FixtureIdentifier,
			Statuses: []*client.Status{
				{Package: "papa", Release: "bookworm"},
				{Package: "papa", Release: "trixie"},
			},
		},
	)
	release := "trixie"
	r, e := m.GetDebianStatus(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		&client.GetDebianStatusParams{Release: &release},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.Statuses)
	all, f := m.GetDebianStatus(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		nil,
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, all.Statuses)
}

func TestMockFlagsLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	b := client.FlagRequest{Identifier: vulnerabilityConstant.FixtureIdentifier}
	_, e := m.FlagVulnerability(x, b)
	assert.FatalOnError(t, e)
	_, f := m.FlagVulnerability(x, b)
	assert.True(t, conflict.Is(f))
	_, g := m.MitigateFlag(x, "CVE-2099-99999", client.MitigationRequest{})
	assert.True(t, not_found.Is(g))
	_, h := m.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{},
	)
	assert.FatalOnError(t, h)
	_, i := m.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{},
	)
	assert.True(t, conflict.Is(i))
	open, j := m.ListFlags(x, nil)
	assert.FatalOnError(t, j)
	assert.Count(t, 0, open)
	reopened, k := m.FlagVulnerability(x, b)
	assert.FatalOnError(t, k)
	assert.False(t, reopened.Mitigated)
	every, l := m.ListFlags(x, nil)
	assert.FatalOnError(t, l)
	assert.Count(t, 1, every)
}
