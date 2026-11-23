package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurity/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestSystemsSendSelection(t *testing.T) {
	o := command_tester.New(t)
	c := o.Fleet()
	names := func(arguments ...string) []string {
		output, _, code := o.Run(
			c,
			append([]string{"systems"}, arguments...)...)
		assert.Integer(t, 0, code)
		var v []*client.SystemSummary
		notation.MustDecode(output, &v, false)
		result := []string{}

		for _, s := range v {
			result = append(result, s.Name)
		}

		return result
	}
	assert.Strings(t, []string{"alfa", "bravo"}, names())
	assert.Strings(t, []string{"alfa"}, names("--label", "group:red"))
	assert.Strings(t, []string{"bravo"}, names("--search", "RAV"))
	assert.Strings(t, []string{}, names("--show", "untracked"))
	assert.Strings(t, []string{}, names("--score", "9.9"))
	_, _, code := o.Run(c, "systems", "--label", "group")
	assert.True(t, code != 0)
}

func TestUpgradesSendSelectionReleaseAndKernel(t *testing.T) {
	o := command_tester.New(t)
	c := o.Fleet()
	sources := func(arguments ...string) []string {
		output, _, code := o.Run(
			c,
			append([]string{"upgrades"}, arguments...)...)
		assert.Integer(t, 0, code)
		var v client.UpgradeList
		notation.MustDecode(output, &v, false)
		result := []string{}

		for _, g := range v.Groups {
			result = append(result, g.Source)
		}

		return result
	}
	assert.Strings(t, []string{"tango-utils"}, sources())
	assert.Strings(t, []string{"linux"}, sources("--kernel"))
	assert.Strings(t, []string{}, sources("--release", "bookworm"))
	assert.Strings(t, []string{}, sources("--label", "group:blue"))
	assert.Strings(t, []string{}, sources("--kernel", "--score", "9"))
}

func TestRemediationDefaultsToCriticalAndSendsSelection(t *testing.T) {
	o := command_tester.New(t)
	c := o.Fleet()
	subjects := func(arguments ...string) []string {
		output, _, code := o.Run(
			c,
			append([]string{"remediation"}, arguments...)...,
		)
		assert.Integer(t, 0, code)
		var v client.RemediationList
		notation.MustDecode(output, &v, false)
		result := []string{}

		for _, i := range v.Items {
			result = append(result, i.Subject)
		}

		return result
	}
	assert.Strings(t, []string{"upgrade:trixie:tango-utils"}, subjects())
	assert.Strings(
		t,
		[]string{"upgrade:trixie:tango-utils", "CVE-2099-00002"},
		subjects("--score", "0"),
	)
	assert.Strings(
		t,
		[]string{"CVE-2099-00002"},
		subjects("--score", "0", "--label", "group:blue"),
	)
}

func TestReadCommandsEmitDaemonResults(t *testing.T) {
	o := command_tester.New(t)
	c := o.Fleet()
	score := 9.8
	c.SeedFindingGroup(
		&client.FindingGroup{Vulnerability: "CVE-2099-00001", Score: &score},
	)
	c.SeedAffectedSystems(&client.AffectedSystems{Identifier: "CVE-2099-00001"})
	c.SeedScoreStatus(client.ScoreStatus{Filled: true, Mirrored: 3})

	for expect, arguments := range map[string][]string{
		`"name":"alfa"`:                       {"system", "alfa"},
		`"vulnerability":"CVE-2099-00001"`:    {"findings", "--score", "9"},
		`"source":"tango-utils"`:              {"upgrade", "trixie", "tango-utils"},
		`"identifier":"CVE-2099-00001"`:       {"affected", "CVE-2099-00001"},
		`"mirrored":3`:                        {"scores"},
	} {
		output, _, code := o.Run(c, arguments...)
		assert.Integer(t, 0, code)
		assert.StringContains(t, expect, output)
	}

	_, _, code := o.Run(c, "system", "oscar")
	assert.True(t, code != 0)
	_, _, empty := o.Run(c, "findings", "--score", "9.9")
	assert.Integer(t, 0, empty)
}

func TestFlagAndMitigateTakeRemediationSubjects(t *testing.T) {
	o := command_tester.New(t)
	c := mock_client.New(t)
	subject := "upgrade:trixie:tango-utils"
	_, _, flagged := o.Run(c, "flag", subject, "--note", "rolling out")
	assert.Integer(t, 0, flagged)
	open, _, code := o.Run(c, "flags")
	assert.Integer(t, 0, code)
	assert.StringContains(t, `"identifier":"upgrade:trixie:tango-utils"`, open)
	_, _, mitigated := o.Run(c, "mitigate", subject, "--how", "rolled out")
	assert.Integer(t, 0, mitigated)
	closed, _, _ := o.Run(c, "flags")
	assert.StringNotContains(t, subject, closed)
	all, _, _ := o.Run(c, "flags", "--all")
	assert.StringContains(t, `"mitigation":"rolled out"`, all)
}
