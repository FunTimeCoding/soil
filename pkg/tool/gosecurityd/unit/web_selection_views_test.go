package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"net/http"
	"testing"
)

func TestUpgradesAndRemediationFollowSelectionOnPagesRestAndTools(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"

	for name, labels := range map[string]map[string]string{
		"alfa":    {"exposure": "internet", "group": "red"},
		"bravo":   {"exposure": "internet", "group": "blue"},
		"charlie": {"exposure": "internal", "group": "red"},
	} {
		assert.FatalOnError(
			t,
			c.PutSystem(
				x,
				name,
				client.SystemSnapshot{
					Release: "trixie",
					Labels:  &labels,
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
	}

	v := s.Served()
	upgrades := v.Body(t, "/upgrades?label=exposure%3Ainternet")
	assert.StringContains(
		t,
		"<strong>2</strong><br><small>Hosts with an upgrade",
		upgrades,
	)
	assert.StringContains(
		t,
		`href="/upgrades?label=exposure%3Ainternet&amp;label=group%3Ared" class="chip">red 1</a>`,
		upgrades,
	)
	assert.StringContains(
		t,
		`href="/upgrades?label=exposure%3Ainternet&amp;release=trixie" class="chip">trixie</a>`,
		upgrades,
	)
	assert.StringContains(
		t,
		`href="/upgrades?kernel=true&amp;label=exposure%3Ainternet" class="chip">Kernel</a>`,
		upgrades,
	)
	kept := v.Body(t, "/upgrades?release=trixie&kernel=true")
	assert.StringContains(
		t,
		`href="/upgrades?kernel=true&amp;label=group%3Ared&amp;release=trixie" class="chip">red 2</a>`,
		kept,
	)
	assert.StringContains(
		t,
		`<input type="hidden" name="release" value="trixie">`,
		kept,
	)

	for _, bad := range []string{
		"/upgrades?kernel=maybe",
		"/upgrades?release=bogus",
		"/upgrades?label=group",
		"/remediation?label=group",
	} {
		assert.Integer(t, http.StatusBadRequest, v.Status(t, bad))
	}

	remediation := v.Body(t, "/remediation?label=group%3Ared&score=0")
	assert.StringContains(
		t,
		`href="/remediation?label=group%3Ared&amp;score=0" class="chip active" aria-current="true">All scores 2</a>`,
		remediation,
	)
	assert.StringContains(
		t,
		`href="/remediation?label=group%3Ared" class="chip">9.0 or higher 0</a>`,
		remediation,
	)
	assert.StringContains(
		t,
		`href="/remediation?label=group%3Ared&amp;score=7" class="chip">7.0 or higher 0</a>`,
		remediation,
	)
	assert.StringContains(
		t,
		`href="/upgrades?release=trixie">tango-utils on trixie</a>`,
		remediation,
	)
	internet := []string{"exposure:internet"}
	listed, e := c.ListUpgrades(x, &client.ListUpgradesParams{Label: &internet})
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, listed.Hosts)
	yes := true
	kernels, i := c.ListUpgrades(x, &client.ListUpgradesParams{Kernel: &yes})
	assert.FatalOnError(t, i)
	assert.Count(t, 0, kernels.Groups)
	everything := 0.0
	red := []string{"group:red"}
	items, f := c.ListRemediation(
		x,
		&client.ListRemediationParams{Label: &red, Score: &everything},
	)
	assert.FatalOnError(t, f)
	assert.Strings(t, []string{"alfa", "charlie"}, items.Items[0].Hosts)
	systems, g := c.ListSystems(x, &client.ListSystemsParams{Label: &internet})
	assert.FatalOnError(t, g)
	assert.Count(t, 2, systems)
	malformed := []string{"group"}
	_, h := c.ListUpgrades(x, &client.ListUpgradesParams{Label: &malformed})
	assert.True(t, validation.Is(h))
	tools := s.Tools()
	internal := map[string]any{"exposure": "internal"}
	toolUpgrades := tools.MustCallToolNotation(
		constant.ListUpgrades,
		map[string]any{constant.LabelsArgument: internal},
	)
	assert.Integer(t, 1, int(toolUpgrades["hosts"].(float64)))
	assert.Count(
		t,
		1,
		tools.MustCallToolNotation(
			constant.ListSystems,
			map[string]any{constant.LabelsArgument: internal},
		)["systems"],
	)
	toolItems := tools.MustCallToolNotation(
		constant.ListRemediation,
		map[string]any{
			constant.LabelsArgument: map[string]any{"group": "blue"},
			constant.ScoreField:     0,
		},
	)["items"].([]any)
	assert.Count(t, 1, toolItems[0].(map[string]any)["hosts"])
	assert.StringContains(
		t,
		"label filter malformed",
		tools.MustCallToolError(
			constant.ListUpgrades,
			map[string]any{constant.LabelsArgument: map[string]any{"group": ""}},
		),
	)
}
