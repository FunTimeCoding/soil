package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestUpgradesGroupByReleaseAndSource(t *testing.T) {
	r, e := security_tester.New(t).
		UpgradeFleet().
		System().
		Upgrades(system_selection.New(), "", false)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, r.Groups)
	first := r.Groups[0]
	assert.String(t, "trixie", first.Release)
	assert.String(t, "tango-utils", first.Source)
	assert.False(t, first.Kernel)
	assert.Count(t, 2, first.Hosts)
	assert.String(t, "bravo", first.Hosts[0].Name)
	assert.String(t, "alfa", first.Hosts[1].Name)
	assert.Strings(
		t,
		[]string{"CVE-2099-00001", "CVE-2099-20002"},
		first.Vulnerabilities,
	)
	assert.String(t, "8.2.0-1", first.InstalledLowest)
	assert.String(t, "8.4.0-1", first.InstalledHighest)
	assert.String(t, "8.6.1-1", first.Candidate)
	assert.String(t, "8.6.1-1", first.Fixed)
	assert.String(t, "bullseye", r.Groups[1].Release)
	assert.Integer(t, 3, r.Hosts)
	assert.Integer(t, 2, r.Vulnerabilities)
	assert.Strings(t, []string{"bullseye", "trixie"}, r.Releases)
}

func TestUpgradesFilterByRelease(t *testing.T) {
	s := security_tester.New(t).UpgradeFleet().System()
	r, e := s.Upgrades(system_selection.New(), "bullseye", false)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.Groups)
	assert.String(t, "charlie", r.Groups[0].Hosts[0].Name)
	_, f := s.Upgrades(system_selection.New(), "bogus", false)
	assert.True(t, validation.Is(f))
}

func TestKernelSidePackagesCarryNoUpgrade(t *testing.T) {
	r, e := security_tester.New(t).
		UpgradeFleet().
		System().
		Upgrades(system_selection.New(), "", true)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Groups)
}

func TestUpgradeListsBinariesPerHost(t *testing.T) {
	s := security_tester.New(t).UpgradeFleet().System()
	g, e := s.Upgrade("trixie", "tango-utils")
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"libtango5"}, g.Hosts[0].Binaries)
	assert.Strings(t, []string{"libtango5", "tango-utils"}, g.Hosts[1].Binaries)
	_, h := s.Upgrade("trixie", "romeo")
	assert.True(t, not_found.Is(h))
}
