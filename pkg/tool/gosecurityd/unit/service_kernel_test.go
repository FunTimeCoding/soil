package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
	"time"
)

func TestKernelMatchesRunningImageOnly(t *testing.T) {
	o := security_tester.New(t).KernelFleet()
	s := o.System()
	d, e := s.Detail("kilo")
	assert.FatalOnError(t, e)
	found := o.KernelFindings(d)
	assert.Count(t, 2, found)
	assert.True(t, found["CVE-2099-60001"].Reboot)
	assert.String(t, "9.1.137-1", found["CVE-2099-60001"].Installed)
	assert.False(t, found["CVE-2099-60003"].Reboot)
	assert.True(t, found["CVE-2099-60003"].Upgradable)
	assert.Integer(t, 1, d.Summary.Reboot)
	assert.Integer(t, 0, d.Summary.Upgradable)
	assert.Integer(t, 0, d.Summary.Findings)
	assert.Integer(t, 2, d.Summary.KernelFindings)
	assert.Integer(t, 2, d.Summary.KernelWithFix)
	assert.String(t, "9.1.137-1", d.Summary.RunningKernel)
	assert.String(t, "9.1.140-1", d.Summary.NewestKernel)
	fixed, f := s.Detail("mike")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, fixed.Findings)
}

func TestBackportsKernelMatchesNextRelease(t *testing.T) {
	o := security_tester.New(t).KernelFleet()
	s := o.System()
	lima, e := s.Detail("lima")
	assert.FatalOnError(t, e)
	found := o.KernelFindings(lima)
	assert.Count(t, 2, found)
	assert.True(t, found["CVE-2099-60001"].Upgradable)
	assert.True(t, found["CVE-2099-60001"].Backports)
	assert.String(t, "9.12.48-1", found["CVE-2099-60001"].Fixed)
	assert.String(t, "no fix", found["CVE-2099-60002"].Fix)
	assert.Integer(t, 2, lima.Summary.KernelFindings)
	assert.Integer(t, 1, lima.Summary.KernelWithFix)
	november, f := s.Detail("november")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, november.Findings)
	assert.String(t, "CVE-2099-60002", november.Findings[0].Vulnerability)
	oscar, g := s.Detail("oscar")
	assert.FatalOnError(t, g)
	other := o.KernelFindings(oscar)
	assert.Count(t, 2, other)
	assert.String(t, "9.1.150-1", other["CVE-2099-60003"].Fixed)
}

func TestKernelWithoutRunningImageHasNoFinding(t *testing.T) {
	d, e := security_tester.New(t).KernelFleet().System().Detail("papa")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, d.Findings)
	assert.String(t, "", d.Summary.RunningKernel)
}

func TestKernelUpgradesListBackportsApart(t *testing.T) {
	s := security_tester.New(t).KernelFleet().System()
	plain, e := s.Upgrades(system_selection.New(), "", false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, plain.Groups)
	r, f := s.Upgrades(system_selection.New(), "", true)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, r.Groups)
	bySource := map[string]int{}

	for i, g := range r.Groups {
		assert.True(t, g.Kernel)
		bySource[g.Source] = i
	}

	release := r.Groups[bySource[constant.KernelSource]]
	assert.String(t, "kilo", release.Hosts[0].Name)
	assert.Strings(t, []string{"CVE-2099-60003"}, release.Vulnerabilities)
	backported := r.Groups[bySource["linux (backports)"]]
	assert.String(t, "lima", backported.Hosts[0].Name)
	g, h := s.Upgrade("bookworm", "linux (backports)")
	assert.FatalOnError(t, h)
	assert.Strings(t, []string{"linux-image-amd64"}, g.Hosts[0].Binaries)
}

func TestKernelKindsCountOnBoardAndHits(t *testing.T) {
	s := security_tester.New(t).KernelFleet().System()
	b, e := s.Board(system_selection.New())
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, b.Kernels[constant.KernelReboot])
	assert.Integer(t, 1, b.Kernels[constant.KernelUpgrade])
	assert.Integer(t, 1, b.Kernels[constant.KernelStuck])
	assert.Integer(t, 2, b.Kernels[constant.KernelCurrent])
	assert.Integer(t, 1, b.Kernels[constant.KernelUnknown])
	names := make([]string, 0, len(b.Rows))

	for _, r := range b.Rows {
		names = append(names, r.Name)
	}

	assert.Strings(
		t,
		[]string{"kilo", "lima", "oscar", "mike", "november", "papa"},
		names,
	)
	hits, f := s.Hits([]string{"CVE-2099-60001"})
	assert.FatalOnError(t, f)
	c := hits["CVE-2099-60001"]
	assert.Integer(t, 3, c.Systems)
	assert.Integer(t, 1, c.Reboot)
	assert.Integer(t, 1, c.Upgrade)
	assert.Integer(t, 1, c.Pending)
	a, g := s.Affected("CVE-2099-60001")
	assert.FatalOnError(t, g)
	assert.String(t, "kilo", a.Rows[0].System)
	assert.String(t, "reboot fixes it", a.Rows[0].Kind)
	assert.String(t, "upgrade and reboot fix it", a.Rows[1].Kind)
}

func TestKernelFollowsOtherTracker(t *testing.T) {
	o := security_tester.New(t).KernelFleet()
	s := o.System()
	before, e := s.Detail("oscar")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, before.Findings)
	o.SecurityTracker.SeedStatus(
		"CVE-2099-60004",
		constant.KernelSource,
		"bookworm",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	o.SecurityTracker.Modify(
		time.Date(2026, time.November, 20, 9, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	after, f := s.Detail("oscar")
	assert.FatalOnError(t, f)
	assert.Count(t, 3, after.Findings)
}
