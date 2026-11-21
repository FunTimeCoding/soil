package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRemediationOnPageRestAndTools(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()

	for _, v := range [][2]string{
		{"CVE-2099-20002", vulnerabilityConstant.DebianResolved},
		{"CVE-2099-20005", vulnerabilityConstant.DebianOpen},
	} {
		fixed := ""

		if v[1] == vulnerabilityConstant.DebianResolved {
			fixed = "8.6.0-1"
		}

		s.SecurityTracker.SeedStatus(v[0], "tango-utils", "trixie", v[1], fixed)
	}

	s.SecurityTracker.SeedUrgency(
		"CVE-2099-20005",
		"tango-utils",
		"trixie",
		constant.Unimportant,
	)
	s.SecurityTracker.SeedStatus(
		"CVE-2099-20008",
		constant.KernelSource,
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"9.12.9-1",
	)

	for identifier, score := range map[string]float64{
		vulnerabilityConstant.FixtureIdentifier: 9.8,
		"CVE-2099-20002":                      7.5,
		"CVE-2099-20005":                      9.5,
		"CVE-2099-20008":                      9.3,
	} {
		s.NationalDatabase.SeedScore(identifier, score, "CRITICAL", time.Now())
	}

	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	candidate := "8.6.1-1"

	for name, offered := range map[string]*string{
		"romeo":  &candidate,
		"sierra": nil,
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
							Version:       "8.4.0-1",
							Source:        "tango-utils",
							SourceVersion: "8.4.0-1",
							Candidate:     offered,
						},
					},
				},
			),
		)
	}

	assert.FatalOnError(
		t,
		c.PutSystem(x, "tango", client.SystemSnapshot{Release: "kilo"}),
	)
	running := "9.12.1-amd64"
	missing := false
	assert.FatalOnError(
		t,
		c.PutSystem(
			x,
			"uniform",
			client.SystemSnapshot{
				Release:                    "trixie",
				Kernel:                     &running,
				KernelMetapackageInstalled: &missing,
				Packages: []*client.InstalledPackage{
					{
						Name:          "linux-image-9.12.1-amd64",
						Version:       "9.12.1-1",
						Source:        "linux-signed-amd64",
						SourceVersion: "9.12.1+1",
					},
				},
			},
		),
	)
	assert.Eventually(
		t,
		func() bool {
			status, e := c.GetScoreStatus(x)
			d, f := c.GetSystemFindings(x, "sierra")

			return e == nil && status.Filled && f == nil && d.Summary.Checked
		},
	)
	note := "rolling out"
	_, e := c.FlagVulnerability(
		x,
		client.FlagRequest{
			Identifier: "upgrade:trixie:tango-utils",
			Note:       &note,
		},
	)
	assert.FatalOnError(t, e)
	_, f := c.FlagVulnerability(
		x,
		client.FlagRequest{Identifier: vulnerabilityConstant.FixtureIdentifier},
	)
	assert.FatalOnError(t, f)
	mitigation := "accepted risk: sierra retires"
	_, g := c.MitigateFlag(
		x,
		vulnerabilityConstant.FixtureIdentifier,
		client.MitigationRequest{Mitigation: &mitigation},
	)
	assert.FatalOnError(t, g)
	list, h := c.ListRemediation(x, nil)
	assert.FatalOnError(t, h)
	order := make([]string, 0, len(list.Items))

	for _, i := range list.Items {
		order = append(order, i.Subject)
	}

	assert.Strings(
		t,
		[]string{
			"kernel:uniform",
			"upgrade:trixie:tango-utils",
			"release:tango",
			"CVE-2099-00001",
		},
		order,
	)
	assert.Integer(t, 3, list.Open)
	assert.Integer(t, 1, list.Addressed)
	assert.Integer(t, 1, list.Unimportant)
	assert.String(t, "kernel metapackage not installed", *list.Items[0].Reason)
	assert.String(t, "rolling out", *list.Items[1].Decision.Note)
	assert.True(t, list.Items[2].Unknown)
	assert.True(t, list.Items[3].Addressed)
	high := 7.0
	wide, i := c.ListRemediation(x, &client.ListRemediationParams{Score: &high})
	assert.FatalOnError(t, i)
	assert.String(t, "CVE-2099-20002", wide.Items[1].Subject)
	assert.Integer(t, 1, wide.Items[1].Count.Pending)
	outside := 11.0
	_, k := c.ListRemediation(x, &client.ListRemediationParams{Score: &outside})
	assert.True(t, validation.Is(k))
	tools := s.Tools()
	assert.Count(
		t,
		5,
		tools.MustCallToolNotation(
			constant.ListRemediation,
			map[string]any{constant.ScoreField: 7},
		)["items"],
	)
	assert.Count(
		t,
		4,
		tools.MustCallToolNotation(constant.ListRemediation, map[string]any{})["items"],
	)
	v := s.Served()
	page := v.Body(t, constant.RemediationPath)
	assert.StringNotContains(t, "CVE-2099-20002", page)
	assert.True(
		t,
		strings.Index(page, "<h2>Addressed by a decision</h2>") <
			strings.Index(page, "Mitigated: </strong>"),
	)
	assert.StringContains(t, "tango-utils on trixie", page)
	assert.StringContains(t, "<code>kernel:uniform</code>", page)
	assert.StringContains(
		t,
		"<small>kernel out of reach: kernel metapackage not installed</small>",
		page,
	)
	assert.StringContains(
		t,
		`href="/system?name=uniform">uniform (trixie)`,
		page,
	)
	assert.StringContains(
		t,
		"<strong>1</strong><br><small>Kernels whose fixes are out of reach",
		page,
	)
	assert.StringContains(t, "<code>upgrade:trixie:tango-utils</code>", page)
	assert.StringContains(t, "Flagged: </strong>rolling out", page)
	assert.StringContains(
		t,
		"Mitigated: </strong>accepted risk: sierra retires",
		page,
	)
	assert.StringContains(t, "<td><small>unknown</small></td>", page)
	assert.StringContains(
		t,
		"<strong>1</strong><br><small>Vulnerabilities left out: Debian marks them unimportant",
		page,
	)
	assert.StringContains(
		t,
		`<a href="/remediation" aria-current="page">Remediation</a>`,
		page,
	)
	assert.Integer(
		t,
		http.StatusBadRequest,
		v.Status(t, "/remediation?score=high"),
	)
	flags := v.Body(t, constant.FlagsPath)
	assert.StringContains(t, "upgrade:trixie:tango-utils", flags)
	assert.StringContains(t, ">upgrade</span>", flags)
	assert.StringContains(
		t,
		`href="/remediation">Remediation`,
		v.Body(t, "/flag/open?identifier=upgrade%3Atrixie%3Atango-utils"),
	)
	assert.StringContains(
		t,
		">unimportant</span>",
		v.Body(t, constant.VulnerabilitiesPath),
	)
	assert.StringContains(
		t,
		"<td><small>unknown</small></td>",
		v.Body(t, constant.SystemsPath),
	)
}
