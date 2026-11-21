package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"net/url"
	"strings"
	"testing"
)

func TestSystemsPageListsAndFiltersSystems(t *testing.T) {
	s := base.New(t)
	s.SecurityTracker.SeedKnown()
	s.Started()
	c := s.SecurityClient()
	candidate := "8.6.1-1"
	assert.FatalOnError(
		t,
		c.PutSystem(
			t.Context(),
			"yankee",
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
	assert.FatalOnError(
		t,
		c.PutSystem(
			t.Context(),
			"zulu",
			client.SystemSnapshot{Release: "lima"},
		),
	)
	assert.Eventually(
		t,
		func() bool {
			d, e := c.GetSystemFindings(t.Context(), "yankee")

			return e == nil && d.Summary.Checked
		},
	)
	v := s.Served()
	all := v.Body(t, constant.SystemsPath)
	assert.StringContains(t, "yankee", all)
	assert.StringContains(t, "zulu", all)
	assert.StringContains(t, "A package upgrade fixes something", all)
	assert.StringContains(t, "Kernel: no upgrade path", all)
	untracked := v.Body(
		t,
		join.Empty(
			constant.SystemsPath,
			"?",
			url.Values{constant.ShowField: {"untracked"}}.Encode(),
		),
	)
	assert.StringContains(t, "zulu", untracked)
	assert.False(t, strings.Contains(untracked, "/system?name=yankee"))
	detail := v.Body(t, "/system?name=yankee")
	assert.StringContains(
		t,
		"Upgrade fixes it: 1 packages, 1 vulnerabilities",
		detail,
	)
	assert.StringContains(t, "tango-utils", detail)
	assert.StringContains(t, "CVE-2099-00001", detail)
}
