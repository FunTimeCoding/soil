package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
	"time"
)

func TestScoresShowOnPagesRestAndTools(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.NationalDatabase.SeedScore(
		vulnerabilityConstant.FixtureIdentifier,
		10.0,
		"CRITICAL",
		time.Now(),
	)
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"
	assert.FatalOnError(
		t,
		c.PutSystem(
			x,
			"romeo",
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
			status, e := c.GetScoreStatus(x)
			d, f := c.GetSystemFindings(x, "romeo")

			return e == nil && status.Filled && f == nil &&
				len(d.Findings) == 1 && d.Findings[0].Score != nil
		},
	)
	d, e := c.GetSystemFindings(x, "romeo")
	assert.FatalOnError(t, e)
	assert.Float(t, 10.0, *d.Findings[0].Score)
	assert.String(t, "CRITICAL", *d.Findings[0].Severity)
	u, f := c.ListUpgrades(x, nil)
	assert.FatalOnError(t, f)
	assert.Float(t, 10.0, u.Groups[0].TopScore)
	a, g := c.GetAffectedSystems(x, vulnerabilityConstant.FixtureIdentifier)
	assert.FatalOnError(t, g)
	assert.Float(t, 10.0, *a.Score)
	status := s.Tools().MustCallToolNotation(
		constant.GetScoreStatus,
		map[string]any{},
	)
	assert.MapValue(t, true, status, "filled")
	v := s.Served()
	assert.StringContains(
		t,
		"Scores: 1 vulnerabilities mirrored from the National Vulnerability Database",
		v.Body(t, constant.SystemsPath),
	)
	assert.StringContains(t, "10.0", v.Body(t, "/system?name=romeo"))
	assert.StringContains(t, "10.0", v.Body(t, constant.UpgradesPath))
}
