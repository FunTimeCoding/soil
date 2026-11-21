package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/kernel_offer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestRemediationSortsEachCriticalFindingIntoOneItem(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	r, e := o.Remediation().List(9)
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{
			"CVE-2099-20006",
			"CVE-2099-20002",
			"upgrade:trixie:tango-utils",
			"kernel:echo",
			"kernel:foxtrot",
			"kernel:hotel",
			"reboot:foxtrot",
			"upgrade:trixie:linux",
			"release:golf",
		},
		o.Subjects(r.Items),
	)
	byKey := map[string]*remediation_item.Item{}

	for _, i := range r.Items {
		byKey[i.Subject] = i
	}

	upgrade := byKey["upgrade:trixie:tango-utils"]
	assert.String(t, "upgrade", upgrade.Kind)
	assert.String(t, "tango-utils", upgrade.Source)
	assert.String(t, "trixie", upgrade.Release)
	assert.Strings(t, []string{"alfa", "bravo"}, upgrade.Hosts)
	assert.Strings(t, []string{"CVE-2099-20002"}, upgrade.Vulnerabilities)
	assert.Float(t, 9.8, upgrade.TopScore)
	pending := byKey["CVE-2099-20002"]
	assert.String(t, "vulnerability", pending.Kind)
	assert.Strings(t, []string{"delta", "echo"}, pending.Hosts)
	assert.Integer(t, 2, pending.Count.Pending)
	assert.String(t, "", pending.Release)
	assert.String(t, "reboot", byKey["reboot:foxtrot"].Kind)
	assert.False(t, upgrade.Kernel)
	kernel := byKey["upgrade:trixie:linux"]
	assert.True(t, kernel.Kernel)
	assert.String(t, "linux", kernel.Source)
	assert.Strings(t, []string{"hotel"}, kernel.Hosts)
	golf := byKey["release:golf"]
	assert.True(t, golf.Unknown)
	assert.String(t, "kilo", golf.Release)
	assert.Count(t, 0, golf.Vulnerabilities)
	assert.Integer(t, 9, r.Open)
	assert.Integer(t, 2, r.Kinds[constant.RemediationVulnerability])
	assert.Integer(t, 2, r.Kinds[constant.RemediationUpgrade])
	assert.Integer(t, 3, r.Kinds[constant.RemediationKernel])
	assert.Integer(t, 1, r.Unimportant)
}

func TestRemediationKernelItemsCarryReason(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	r, e := o.Remediation().List(9)
	assert.FatalOnError(t, e)
	byKey := map[string]*remediation_item.Item{}

	for _, i := range r.Items {
		byKey[i.Subject] = i
	}

	echo := byKey["kernel:echo"]
	assert.String(t, "kernel", echo.Kind)
	assert.String(t, "trixie", echo.Release)
	assert.Strings(t, []string{"echo"}, echo.Hosts)
	assert.Strings(
		t,
		[]string{"CVE-2099-20001", "CVE-2099-20007"},
		echo.Vulnerabilities,
	)
	assert.String(t, "sources not reported yet", echo.Reason)
	hotel := byKey["kernel:hotel"]
	assert.Strings(t, []string{"CVE-2099-20007"}, hotel.Vulnerabilities)
	assert.String(t, "sources offer up to 9.12.2-1", hotel.Reason)
	assert.Strings(
		t,
		[]string{"CVE-2099-20007"},
		byKey["kernel:foxtrot"].Vulnerabilities,
	)
	assert.Strings(
		t,
		[]string{"CVE-2099-20001"},
		byKey["reboot:foxtrot"].Vulnerabilities,
	)
	unfixed := byKey["CVE-2099-20006"]
	assert.String(t, "vulnerability", unfixed.Kind)
	assert.Strings(t, []string{"echo", "foxtrot", "hotel"}, unfixed.Hosts)
	assert.Integer(t, 3, unfixed.Count.Open)
	d, f := o.System().Detail("hotel")
	assert.FatalOnError(t, f)
	assert.String(t, "upgrade and reboot", d.Summary.KernelState)
	assert.String(t, "sources offer up to 9.12.2-1", d.Summary.KernelReason)
}

func TestMissingMetapackageNamesKernelReason(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	s := o.System()
	missing := false

	for name, offers := range map[string][]*kernel_offer.Offer{
		"india":   {kernel_offer.New("9.12.9-1", 500, []string{"trixie-security"})},
		"juliett": nil,
	} {
		v := system.New(name, "trixie", 0)
		v.Kernel = "9.12.1-amd64"
		v.KernelOffers = offers
		v.KernelMetapackageInstalled = &missing
		assert.FatalOnError(
			t,
			s.Put(
				v,
				[]installed.Installed{
					*installed.New(
						"linux-image-9.12.1-amd64",
						"9.12.1-1",
						"linux-signed-amd64",
						"9.12.1+1",
						"",
					),
				},
			),
		)
	}

	india, e := s.Detail("india")
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"kernel metapackage not installed (sources offer up to 9.12.9-1)",
		india.Summary.KernelReason,
	)
	juliett, f := s.Detail("juliett")
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"kernel metapackage not installed",
		juliett.Summary.KernelReason,
	)
	r, g := o.Remediation().List(9)
	assert.FatalOnError(t, g)

	for _, i := range r.Items {
		if i.Subject == "kernel:india" {
			assert.String(
				t,
				"kernel metapackage not installed (sources offer up to 9.12.9-1)",
				i.Reason,
			)
		}
	}
}

func TestRemediationRanksByHostsTimesVulnerabilities(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	r, e := o.Remediation().List(7)
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{
			"upgrade:trixie:tango-utils",
			"CVE-2099-00001",
			"CVE-2099-20006",
			"CVE-2099-20002",
			"kernel:echo",
			"kernel:foxtrot",
			"kernel:hotel",
			"reboot:foxtrot",
			"release:charlie",
			"upgrade:trixie:linux",
			"release:golf",
		},
		o.Subjects(r.Items),
	)
	known := r.Items[1]
	assert.Strings(t, []string{"alfa", "delta", "echo"}, known.Hosts)
	assert.Integer(t, 2, known.Count.Open)
	assert.Integer(t, 1, known.Count.Pending)
	charlie := r.Items[8]
	assert.String(t, "bullseye", charlie.Release)
	assert.False(t, charlie.Unknown)
	assert.Strings(t, []string{"CVE-2099-00001"}, charlie.Vulnerabilities)
}

func TestRemediationJoinsDecisions(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	x := t.Context()
	_, e := o.Flag().Flag(x, "upgrade:trixie:tango-utils", "rolling out")
	assert.FatalOnError(t, e)
	_, f := o.Flag().Flag(x, "reboot:foxtrot", "")
	assert.FatalOnError(t, f)
	_, g := o.Flag().Mitigate("reboot:foxtrot", "accepted: replaced next week")
	assert.FatalOnError(t, g)
	r, h := o.Remediation().List(9)
	assert.FatalOnError(t, h)
	assert.Integer(t, 8, r.Open)
	assert.Integer(t, 1, r.Addressed)
	assert.Integer(t, 0, r.Kinds[constant.RemediationReboot])
	last := r.Items[len(r.Items)-1]
	assert.String(t, "reboot:foxtrot", last.Subject)
	assert.True(t, last.Addressed)
	assert.String(t, "accepted: replaced next week", last.Decision.Mitigation)

	for _, i := range r.Items {
		if i.Subject == "upgrade:trixie:tango-utils" {
			assert.String(t, "rolling out", i.Decision.Note)
			assert.False(t, i.Addressed)
		}
	}
}

func TestRemediationItemLeavesWhenNoHostNeedsIt(t *testing.T) {
	o := security_tester.New(t).RemediationFleet()
	o.PutFixturePackage("delta", "trixie", "8.6.1-1", "")
	o.PutFixturePackage("echo", "trixie", "8.6.1-1", "")
	r, e := o.Remediation().List(9)
	assert.FatalOnError(t, e)

	for _, i := range r.Items {
		assert.True(t, i.Subject != "CVE-2099-20002")
	}
}

func TestRemediationScoreOutsideRangeIsValidation(t *testing.T) {
	_, e := security_tester.New(t).Remediation().List(11)
	assert.True(t, validation.Is(e))
}

func TestGroupsMarkUnimportant(t *testing.T) {
	groups, e := security_tester.New(t).RemediationFleet().System().Groups(9)
	assert.FatalOnError(t, e)

	for _, g := range groups {
		assert.Boolean(t, g.Vulnerability == "CVE-2099-20005", g.Unimportant)
	}
}
