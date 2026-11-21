package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestMalformedIdentifierReachesClientAsValidation(t *testing.T) {
	c := base.New(t).SecurityClient()
	_, e := c.GetVulnerability(t.Context(), "november")
	assert.True(t, validation.Is(e))
	_, f := c.GetDebianStatus(t.Context(), "november", nil)
	assert.True(t, validation.Is(f))
	_, g := c.FlagVulnerability(
		t.Context(),
		client.FlagRequest{Identifier: "november"},
	)
	assert.True(t, validation.Is(g))
	_, h := c.MitigateFlag(t.Context(), "november", client.MitigationRequest{})
	assert.True(t, validation.Is(h))
}

func TestUnknownVulnerabilityReachesClientAsNotFound(t *testing.T) {
	_, e := base.New(t).SecurityClient().GetVulnerability(
		t.Context(),
		"CVE-2099-99999",
	)
	assert.True(t, not_found.Is(e))
}

func TestUnloadedTrackerReachesClientAsUnavailable(t *testing.T) {
	_, e := base.New(t).SecurityClient().GetDebianStatus(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		nil,
	)
	assert.True(t, unavailable.Is(e))
}

func TestFlagRefusalsReachClientByClass(t *testing.T) {
	s := base.New(t)
	c := s.SecurityClient()
	x := t.Context()
	_, e := c.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{},
	)
	assert.True(t, not_found.Is(e))
	b := client.FlagRequest{Identifier: vulnerabilityConstant.FixtureIdentifier}
	_, f := c.FlagVulnerability(x, b)
	assert.FatalOnError(t, f)
	_, g := c.FlagVulnerability(x, b)
	assert.True(t, conflict.Is(g))
	_, h := c.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{},
	)
	assert.FatalOnError(t, h)
	_, i := c.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{},
	)
	assert.True(t, conflict.Is(i))
	assert.Count(t, 0, s.Reporter.Events())
}
