package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestUpgradesBuildFromSelectedSystemsOnly(t *testing.T) {
	s := security_tester.New(t).LabeledFleet().System()
	internet, e := s.Upgrades(
		system_selection.New().WithLabel("exposure", "internet"),
		"",
		false,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, internet.Groups)
	names := []string{}

	for _, h := range internet.Groups[0].Hosts {
		names = append(names, h.Name)
	}

	assert.Strings(t, []string{"alfa", "bravo"}, names)
	assert.Integer(t, 2, internet.Hosts)
	found, f := s.Upgrades(system_selection.New().WithSearch("CHAR"), "", false)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, found.Hosts)
	assert.String(t, "charlie", found.Groups[0].Hosts[0].Name)
	untracked, g := s.Upgrades(
		system_selection.New().WithCoverage(constant.ShowUntracked),
		"",
		false,
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, untracked.Groups)
}

func TestRemediationBuildsFromSelectedSystemsOnly(t *testing.T) {
	o := security_tester.New(t).LabeledFleet()
	echo := system.New("echo", "kilo", 0)
	echo.Labels = map[string]string{"exposure": "internet"}
	assert.FatalOnError(t, o.System().Put(echo, nil))
	internet := system_selection.New().WithLabel("exposure", "internet")
	r, e := o.Remediation().List(internet.WithMinimum(0))
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"upgrade:trixie:tango-utils", "release:echo"},
		o.Subjects(r.Items),
	)
	assert.Strings(t, []string{"alfa", "bravo"}, r.Items[0].Hosts)
	critical, f := o.Remediation().List(internet.WithMinimum(9))
	assert.FatalOnError(t, f)
	assert.Strings(t, []string{"release:echo"}, o.Subjects(critical.Items))
	internal, g := o.Remediation().List(
		system_selection.New().WithLabel("exposure", "internal"),
	)
	assert.FatalOnError(t, g)
	assert.Strings(
		t,
		[]string{"upgrade:trixie:tango-utils"},
		o.Subjects(internal.Items),
	)
	assert.Strings(t, []string{"charlie", "delta"}, internal.Items[0].Hosts)
}
