package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"net/url"
	"testing"
)

func TestUpgradesAnswerOnPageRestAndTools(t *testing.T) {
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
			"echo",
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
					{
						Name:          "tango-utils",
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
			d, e := c.GetSystemFindings(x, "echo")

			return e == nil && d.Summary.Checked
		},
	)
	list, e := c.ListUpgrades(x, nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, list.Groups)
	assert.Integer(t, 1, list.Hosts)
	bogus := "bogus"
	_, f := c.ListUpgrades(x, &client.ListUpgradesParams{Show: &bogus})
	assert.True(t, validation.Is(f))
	g, h := c.GetUpgrade(x, "trixie", "tango-utils")
	assert.FatalOnError(t, h)
	assert.Strings(
		t,
		[]string{"libtango5", "tango-utils"},
		*g.Hosts[0].Binaries,
	)
	_, k := c.GetUpgrade(x, "trixie", "romeo")
	assert.True(t, not_found.Is(k))
	v := s.Served()
	page := v.Body(t, constant.UpgradesPath)
	assert.StringContains(t, "tango-utils", page)
	assert.StringContains(t, "Package upgrades", page)
	query := url.Values{
		constant.ReleaseParameter: {"trixie"},
		constant.SourceParameter:  {"tango-utils"},
	}.Encode()
	opened := v.Body(t, join.Empty(constant.UpgradeOpenPath, "?", query))
	assert.StringContains(
		t,
		"apt-get install --only-upgrade libtango5 tango-utils",
		opened,
	)
	assert.StringContains(t, "/system?name=echo", opened)
	closed := v.Body(t, join.Empty(constant.UpgradeShutPath, "?", query))
	assert.StringContains(t, "upgrade-trixie-tango-utils", closed)
	tools := s.Tools()
	listed := tools.MustCallToolNotation(
		constant.ListUpgrades,
		map[string]any{},
	)
	assert.Count(t, 1, listed["groups"])
	upgrade := tools.MustCallToolNotation(
		constant.GetUpgrade,
		map[string]any{
			constant.ReleaseParameter: "trixie",
			constant.SourceParameter:  "tango-utils",
		},
	)
	assert.MapValue(t, "tango-utils", upgrade, constant.SourceParameter)
}
