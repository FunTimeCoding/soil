package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/probe"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
	"time"
)

func TestCodenameOutputReadsPackagesAndCandidates(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("codename"))
	assert.FatalOnError(t, e)
	assert.String(t, "trixie", s.Release)
	assert.String(t, "9.12.48+deb13-amd64", *s.Kernel)
	assert.Count(t, 4, s.Packages)
	byName := map[string]string{}

	for _, p := range s.Packages {
		assert.True(t, p.Source != "" && p.SourceVersion != "")

		if p.Candidate != nil {
			byName[p.Name] = *p.Candidate
		}
	}

	assert.String(t, "8.6.1-1", byName["libtango5"])
	assert.String(t, "2.91-13", byName["libuniform6"])
	assert.String(t, "2.91-14", byName["libuniform6:i386"])
	_, upgradable := byName["tango-utils"]
	assert.False(t, upgradable)
}

func TestListsDateAndKernelOffersAreRead(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("codename"))
	assert.FatalOnError(t, e)
	assert.Time(t, time.Unix(1763538000, 0).UTC(), *s.ListsUpdated)
	offers := *s.KernelOffers
	assert.Count(t, 2, offers)
	assert.String(t, "9.12.57-1", offers[0].Version)
	assert.Integer(t, 500, offers[0].Priority)
	assert.Strings(
		t,
		[]string{"http://deb.debian.org/debian-security trixie-security/main"},
		offers[0].Origins,
	)
	assert.String(t, "9.12.48-1", offers[1].Version)
	assert.Strings(
		t,
		[]string{
			"http://deb.debian.org/debian trixie/main",
			"/var/lib/dpkg/status",
		},
		offers[1].Origins,
	)
}

func TestMetapackageStateComesFromInstalledLine(t *testing.T) {
	missing, e := probe.Parse(
		security_tester.ProbeFixture("metapackage-removed"),
	)
	assert.FatalOnError(t, e)
	assert.False(t, *missing.KernelMetapackageInstalled)
	assert.Strings(
		t,
		[]string{"/var/lib/dpkg/status"},
		(*missing.KernelOffers)[1].Origins,
	)
	present, f := probe.Parse(security_tester.ProbeFixture("codename"))
	assert.FatalOnError(t, f)
	assert.True(t, *present.KernelMetapackageInstalled)
	unreported, g := probe.Parse(
		security_tester.ProbeFixture("red-hat-release"),
	)
	assert.FatalOnError(t, g)
	assert.Nil(t, unreported.KernelMetapackageInstalled)
}

func TestMissingListsAndPolicyLeaveFieldsEmpty(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("red-hat-release"))
	assert.FatalOnError(t, e)
	assert.Nil(t, s.ListsUpdated)
	assert.Nil(t, s.KernelOffers)
	junk, f := probe.Parse(
		"##release\nID=debian\nVERSION_CODENAME=trixie\n##lists\nsierra\n",
	)
	assert.FatalOnError(t, f)
	assert.Nil(t, junk.ListsUpdated)
	stray, g := probe.Parse(
		"##release\nID=debian\nVERSION_CODENAME=trixie\n##kernel-policy\n        500 http://deb.debian.org/debian trixie/main amd64 Packages\n",
	)
	assert.FatalOnError(t, g)
	assert.Nil(t, stray.KernelOffers)
}

func TestReleaseWithoutCodenameComesFromVersion(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("without-codename"))
	assert.FatalOnError(t, e)
	assert.String(t, "kilo", s.Release)
	assert.Count(t, 1, s.Packages)
}

func TestOtherReleaseComesFromIdentifierAndMajor(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("other-identifier"))
	assert.FatalOnError(t, e)
	assert.String(t, "november7", s.Release)
	assert.Count(t, 0, s.Packages)
}

func TestReleaseWithoutOperatingSystemReleaseReadsRedHatRelease(t *testing.T) {
	s, e := probe.Parse(security_tester.ProbeFixture("red-hat-release"))
	assert.FatalOnError(t, e)
	assert.String(t, "mike6", s.Release)
	assert.String(t, "9.6.0-1.mike6.x86_64", *s.Kernel)
}

func TestOutputWithoutReleaseIsRefused(t *testing.T) {
	_, e := probe.Parse("host did not answer")
	assert.Error(t, e)
}

func TestUnknownReleaseIsRefused(t *testing.T) {
	_, e := probe.Parse("##release\nNAME=sierra\n##kernel\n1.0\n")
	assert.Error(t, e)
}
