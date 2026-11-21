package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/url"
	"testing"
)

func TestAffectedSystemsAnswerOnPagesRestAndTools(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"
	assert.FatalOnError(
		t,
		c.PutSystem(
			x,
			"hotel",
			client.SystemSnapshot{
				Release: "trixie",
				Packages: []*client.InstalledPackage{
					{
						Name:          "libtango5",
						Version:       "8.4.0-1",
						Source:        "tango-utils",
						SourceVersion: "8.4.0-1",
						Candidate:     &candidate,
					},
				},
			},
		),
	)
	assert.Eventually(
		t,
		func() bool {
			d, e := c.GetSystemFindings(x, "hotel")

			return e == nil && d.Summary.Checked
		},
	)
	a, e := c.GetAffectedSystems(x, vulnerabilityConstant.FixtureIdentifier)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, a.Count.Upgrade)
	assert.String(t, "hotel", a.Systems[0].System)
	_, f := c.GetAffectedSystems(x, "bogus")
	assert.True(t, validation.Is(f))
	tools := s.Tools()
	affected := tools.MustCallToolNotation(
		constant.GetAffectedSystems,
		map[string]any{"identifier": vulnerabilityConstant.FixtureIdentifier},
	)
	assert.Count(t, 1, affected["systems"])
	v := s.Served()
	query := url.Values{
		"identifier": {vulnerabilityConstant.FixtureIdentifier},
	}.Encode()
	lookup := v.Body(t, join.Empty(constant.LookupPath, "?", query))
	assert.StringContains(t, "Affected systems", lookup)
	assert.StringContains(t, "/system?name=hotel", lookup)
	assert.StringContains(t, "1 systems: an upgrade fixes 1", lookup)
	v.PostBody(
		t,
		constant.FlagPath,
		url.Values{"identifier": {vulnerabilityConstant.FixtureIdentifier}},
	)
	flags := v.Body(t, constant.FlagsPath)
	assert.StringContains(t, "1 upgrade fixes", flags)
	opened := v.Body(t, join.Empty(constant.FlagOpenPath, "?", query))
	assert.StringContains(t, "Affected systems: ", opened)
	assert.StringContains(t, "1 systems: an upgrade fixes 1", opened)
	assert.StringContains(t, "#affected-systems", opened)
}
