package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"net/url"
	"testing"
	"time"
)

func TestKernelShowsOnPages(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedStatus(
		"CVE-2099-70001",
		constant.KernelSource,
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"9.12.48-1",
	)
	s.Started()
	c := s.SecurityClient()
	x := t.Context()
	kernel := "9.12.38+deb12-amd64"
	candidate := "9.12.95-1~bpo12+1"
	lists := time.Date(2026, time.November, 19, 6, 0, 0, 0, time.UTC)
	offers := []client.KernelOffer{
		{
			Version:  "9.12.95-1~bpo12+1",
			Priority: 100,
			Origins: []string{
				"http://deb.debian.org/debian bookworm-backports/main",
			},
		},
	}
	assert.FatalOnError(
		t,
		c.PutSystem(
			x,
			"quebec",
			client.SystemSnapshot{
				Release:      "bookworm",
				Kernel:       &kernel,
				ListsUpdated: &lists,
				KernelOffers: &offers,
				Packages: []*client.InstalledPackage{
					{
						Name:          "linux-image-9.12.38+deb12-amd64",
						Version:       "9.12.38-1~bpo12+1",
						Source:        "linux-signed-amd64",
						SourceVersion: "9.12.38+1~bpo12+1",
					},
					{
						Name:          "linux-image-amd64",
						Version:       "9.12.38-1~bpo12+1",
						Source:        "linux-signed-amd64",
						SourceVersion: "9.12.38+1~bpo12+1",
						Candidate:     &candidate,
					},
				},
			},
		),
	)
	assert.Eventually(
		t,
		func() bool {
			d, e := c.GetSystemFindings(x, "quebec")

			return e == nil && d.Summary.Checked && len(d.Findings) == 1
		},
	)
	d, e := c.GetSystemFindings(x, "quebec")
	assert.FatalOnError(t, e)
	assert.String(t, "upgrade and reboot", d.Summary.KernelState)
	assert.Integer(t, 1, d.Summary.KernelWithFix)
	assert.Time(t, lists, *d.Summary.ListsUpdated)
	assert.Count(t, 1, *d.Summary.KernelOffers)
	v := s.Served()
	detail := v.Body(t, "/system?name=quebec")
	assert.StringContains(t, "9.12.38-1~bpo12+1", detail)
	assert.StringContains(t, "bookworm-backports/main", detail)
	assert.StringContains(
		t,
		"Kernel: 1 vulnerabilities on the running kernel",
		detail,
	)
	board := v.Body(t, constant.SystemsPath)
	assert.StringContains(t, "upgrade and reboot", board)
	query := url.Values{
		constant.ReleaseParameter: {"bookworm"},
		constant.SourceParameter:  {"linux (backports)"},
	}.Encode()
	opened := v.Body(t, join.Empty(constant.UpgradeOpenPath, "?", query))
	assert.StringContains(t, "--only-upgrade linux-image-amd64", opened)
	assert.StringContains(t, "Then reboot into the new kernel", opened)
}
