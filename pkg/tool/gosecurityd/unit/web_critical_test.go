package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/http"
	"testing"
	"time"
)

func TestScoreFilterOnPagesRestAndTools(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.SecurityTracker.SeedStatus(
		"CVE-2099-20002",
		"tango-utils",
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"8.6.0-1",
	)
	s.NationalDatabase.SeedScore(
		vulnerabilityConstant.FixtureIdentifier,
		5.0,
		"MEDIUM",
		time.Now(),
	)
	s.NationalDatabase.SeedScore("CVE-2099-20002", 9.8, "CRITICAL", time.Now())
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"

	for name, version := range map[string]string{
		"romeo":  "8.4.0-1",
		"sierra": "8.6.0-1",
	} {
		assert.FatalOnError(
			t,
			c.PutSystem(
				x,
				name,
				client.SystemSnapshot{
					Release: "trixie",
					Packages: []*client.InstalledPackage{
						{
							Name:          "libtango5",
							Version:       version,
							Source:        "tango-utils",
							SourceVersion: version,
							Candidate:     &candidate,
						},
					},
				},
			),
		)
	}

	assert.Eventually(
		t,
		func() bool {
			status, e := c.GetScoreStatus(x)
			d, f := c.GetSystemFindings(x, "romeo")

			return e == nil && status.Filled && f == nil && d.Summary.Checked
		},
	)
	critical := 9.0
	narrowed, e := c.ListSystems(x, &client.ListSystemsParams{Score: &critical})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, narrowed)
	assert.String(t, "romeo", narrowed[0].Name)
	assert.Integer(t, 1, narrowed[0].Critical)
	assert.Float(t, 9.8, narrowed[0].TopScore)
	groups, f := c.ListFindings(x, &client.ListFindingsParams{Score: &critical})
	assert.FatalOnError(t, f)
	assert.Count(t, 1, groups)
	assert.Float(t, 9.8, *groups[0].Score)
	assert.Integer(t, 1, groups[0].Count.Upgrade)
	all, g := c.ListFindings(x, nil)
	assert.FatalOnError(t, g)
	assert.Count(t, 2, all)
	upgrades, h := c.ListUpgrades(
		x,
		&client.ListUpgradesParams{Score: &critical},
	)
	assert.FatalOnError(t, h)
	assert.Strings(
		t,
		[]string{"CVE-2099-20002"},
		upgrades.Groups[0].Vulnerabilities,
	)
	outside := 11.0
	_, i := c.ListSystems(x, &client.ListSystemsParams{Score: &outside})
	assert.True(t, validation.Is(i))
	_, k := c.ListFindings(x, &client.ListFindingsParams{Score: &outside})
	assert.True(t, validation.Is(k))
	tools := s.Tools()
	assert.Count(
		t,
		1,
		tools.MustCallToolNotation(
			constant.ListFindings,
			map[string]any{constant.ScoreField: 9},
		)["findings"],
	)
	assert.Count(
		t,
		1,
		tools.MustCallToolNotation(
			constant.ListSystems,
			map[string]any{constant.ScoreField: 9},
		)["systems"],
	)
	toolUpgrades := tools.MustCallToolNotation(
		constant.ListUpgrades,
		map[string]any{constant.ScoreField: 9},
	)["groups"].([]any)
	assert.Count(t, 1, toolUpgrades)
	assert.Count(t, 1, toolUpgrades[0].(map[string]any)["vulnerabilities"])
	v := s.Served()
	opening := v.Body(t, constant.VulnerabilitiesPath)
	assert.StringContains(t, "CVE-2099-20002", opening)
	assert.StringNotContains(t, "CVE-2099-00001", opening)
	assert.StringContains(
		t,
		`href="/vulnerabilities?score=9" class="chip active"`,
		opening,
	)
	assert.StringContains(t, `href="/vulnerabilities?score=0"`, opening)
	assert.StringContains(
		t,
		"<td>1</td><td>0</td><td>1</td><td>0</td><td>0</td><td><small>0</small></td>",
		opening,
	)
	assert.StringContains(
		t,
		"<strong>1</strong><br><small>An upgrade or reboot fixes it somewhere",
		opening,
	)
	assert.StringContains(
		t,
		`<a href="/vulnerabilities" aria-current="page">Vulnerabilities</a>`,
		opening,
	)
	assert.StringContains(
		t,
		"CVE-2099-00001",
		v.Body(t, "/vulnerabilities?score=0"),
	)
	assert.Integer(
		t,
		http.StatusBadRequest,
		v.Status(t, "/vulnerabilities?score=high"),
	)
	assert.Integer(t, http.StatusBadRequest, v.Status(t, "/systems?score=11"))
	systems := v.Body(t, "/systems?score=9")
	assert.StringContains(t, "romeo", systems)
	assert.StringNotContains(t, "sierra", systems)
	assert.StringContains(
		t,
		"<strong>1</strong><br><small>A vulnerability at 9.0 or higher",
		systems,
	)
	assert.StringContains(t, `--badge-color: var(--bad)">1</span>`, systems)
	assert.StringContains(
		t,
		`href="/systems?score=9&amp;show=supported"`,
		systems,
	)
	assert.StringContains(
		t,
		`href="/systems?score=9" class="chip active" aria-current="true">9.0 or higher`,
		systems,
	)
	assert.StringContains(t, "sierra", v.Body(t, constant.SystemsPath))
	page := v.Body(t, "/system?name=romeo&score=9")
	assert.StringContains(t, "CVE-2099-20002", page)
	assert.StringNotContains(t, "CVE-2099-00001", page)
	assert.StringContains(
		t,
		`href="/system?name=romeo&amp;score=9" class="chip active"`,
		page,
	)
	fixes := "<strong>%d</strong><br><small>Vulnerabilities an upgrade fixes"
	narrowedUpgrades := v.Body(t, "/upgrades?score=9")
	assert.StringContains(t, fmt.Sprintf(fixes, 1), narrowedUpgrades)
	assert.StringContains(
		t,
		`href="/upgrades?release=trixie&amp;score=9"`,
		narrowedUpgrades,
	)
	assert.StringContains(
		t,
		fmt.Sprintf(fixes, 2),
		v.Body(t, constant.UpgradesPath),
	)
}
