package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"testing"
)

func TestClientRoundTripsSystemSnapshot(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "5.6.1-1"
	priority := 1
	assert.FatalOnError(
		t,
		c.PutSystem(
			x,
			"kilo",
			client.SystemSnapshot{
				Release:  "trixie",
				Priority: &priority,
				Labels:   &map[string]string{"exposure": "internet"},
				Packages: []*client.InstalledPackage{
					{
						Name:          "liblzma5",
						Version:       "5.4.0-1",
						Source:        "xz-utils",
						SourceVersion: "5.4.0-1",
						Candidate:     &candidate,
					},
				},
			},
		),
	)
	var d *client.SystemDetail
	assert.Eventually(
		t,
		func() bool {
			var e error
			d, e = c.GetSystemFindings(x, "kilo")

			return e == nil && d.Summary.Checked
		},
	)
	assert.String(t, "supported", d.Summary.Coverage)
	assert.False(t, d.Summary.Stale)
	assert.False(t, d.Summary.TrackerStale)
	assert.NotNil(t, d.Summary.TrackerVerified)
	assert.False(t, d.Summary.Clean)
	assert.Count(t, 1, d.Findings)
	assert.String(t, "debian", d.Findings[0].Fix)
	assert.True(t, d.Findings[0].Upgradable)
	assert.String(t, "internet", (*d.Summary.Labels)["exposure"])
	systems, f := c.ListSystems(x)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, systems)
	assert.Integer(t, 1, systems[0].Findings)
	groups, g := c.ListFindings(x)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, groups)
	assert.Strings(t, []string{"kilo"}, groups[0].Systems)
	assert.FatalOnError(t, c.DeleteSystem(x, "kilo"))
	_, h := c.GetSystemFindings(x, "kilo")
	assert.True(t, not_found.Is(h))
	assert.True(t, not_found.Is(c.DeleteSystem(x, "kilo")))
}

func TestMalformedSnapshotReachesClientAsValidation(t *testing.T) {
	c := base.New(t).SecurityClient()
	e := c.PutSystem(t.Context(), "lima", client.SystemSnapshot{Release: ""})
	assert.True(t, validation.Is(e))
}

func TestUntrackedSystemOverRestIsNotClean(t *testing.T) {
	s := base.New(t)
	s.Started()
	c := s.SecurityClient()
	assert.FatalOnError(
		t,
		c.PutSystem(
			t.Context(),
			"mike",
			client.SystemSnapshot{Release: "wheezy"},
		),
	)
	d, e := c.GetSystemFindings(t.Context(), "mike")
	assert.FatalOnError(t, e)
	assert.String(t, "unsupported, untracked", d.Summary.Coverage)
	assert.False(t, d.Summary.Checked)
	assert.False(t, d.Summary.Clean)
}

func TestSystemToolsAnswer(t *testing.T) {
	s := base.New(t)
	s.Started()
	assert.FatalOnError(
		t,
		s.SecurityClient().PutSystem(
			t.Context(),
			"november",
			client.SystemSnapshot{
				Release: "bullseye",
				Packages: []*client.InstalledPackage{
					{
						Name:          "liblzma5",
						Version:       "5.2.5-2.1~deb11u1",
						Source:        "xz-utils",
						SourceVersion: "5.2.5-2.1~deb11u1",
					},
				},
			},
		),
	)
	tools := s.Tools()
	systems := tools.MustCallToolNotation(
		constant.ListSystems,
		map[string]any{},
	)
	assert.Count(t, 1, systems["systems"])
	detail := tools.MustCallToolNotation(
		constant.GetSystemFindings,
		map[string]any{"name": "november"},
	)
	summary := detail["summary"].(map[string]any)
	assert.MapValue(t, "unsupported, tracked", summary, "coverage")
	assert.MapValue(t, false, summary, "clean")
	findings := tools.MustCallToolNotation(
		constant.ListFindings,
		map[string]any{},
	)
	assert.Count(t, 0, findings["findings"])
	assert.StringContains(
		t,
		"system not found: oscar",
		tools.MustCallToolError(
			constant.GetSystemFindings,
			map[string]any{"name": "oscar"},
		),
	)
}
