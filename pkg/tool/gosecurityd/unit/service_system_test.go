package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
	"time"
)

func TestSupportedSystemWithFixedPackageIsClean(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("tango", "trixie", "5.6.1-1", "")
	d, e := o.System().Detail("tango")
	assert.FatalOnError(t, e)
	assert.String(t, "supported", d.Summary.Coverage)
	assert.True(t, d.Summary.Checked)
	assert.False(t, d.Summary.Stale)
	assert.False(t, d.Summary.TrackerStale)
	assert.True(t, d.Summary.Clean)
	assert.Count(t, 0, d.Findings)
}

func TestStaleSnapshotIsNotClean(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("oscar", "trixie", "5.6.1-1", "")
	o.Later = constant.StaleAfter + time.Hour
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	d, e := o.System().Detail("oscar")
	assert.FatalOnError(t, e)
	assert.True(t, d.Summary.Checked)
	assert.True(t, d.Summary.Stale)
	assert.False(t, d.Summary.TrackerStale)
	assert.Count(t, 0, d.Findings)
	assert.False(t, d.Summary.Clean)
}

func TestStaleTrackerIsNotClean(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("papa", "trixie", "5.6.1-1", "")
	o.Later = constant.TrackerStaleAfter + time.Hour
	d, e := o.System().Detail("papa")
	assert.FatalOnError(t, e)
	assert.True(t, d.Summary.Checked)
	assert.False(t, d.Summary.Stale)
	assert.True(t, d.Summary.TrackerStale)
	assert.Count(t, 0, d.Findings)
	assert.False(t, d.Summary.Clean)
}

func TestStaleSnapshotOnUntrackedReleaseIsMarked(t *testing.T) {
	o := security_tester.New(t)
	o.PutXz("quebec", "wheezy", "5.1.1alpha+20120614-2", "")
	o.Later = constant.StaleAfter + time.Hour
	d, e := o.System().Detail("quebec")
	assert.FatalOnError(t, e)
	assert.True(t, d.Summary.Stale)
	assert.False(t, d.Summary.TrackerStale)
	assert.True(t, d.Summary.TrackerVerified.IsZero())
}

func TestOlderPackageOnSupportedSystemIsFinding(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("uniform", "trixie", "5.4.0-1", "5.6.1-1")
	d, e := o.System().Detail("uniform")
	assert.FatalOnError(t, e)
	assert.False(t, d.Summary.Clean)
	assert.Count(t, 1, d.Findings)
	f := d.Findings[0]
	assert.String(t, "CVE-2024-3094", f.Vulnerability)
	assert.String(t, "debian", f.Fix)
	assert.String(t, "5.6.1-1", f.Fixed)
	assert.String(t, "5.4.0-1", f.Installed)
	assert.True(t, f.Upgradable)
	assert.Integer(t, 1, d.Summary.Upgradable)
}

func TestOpenStatusIsFindingWithoutFix(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedStatus(
		vulnerabilityConstant.KnownIdentifier,
		"xz-utils",
		"trixie",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("victor", "trixie", "5.8.1-1", "")
	d, e := o.System().Detail("victor")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, d.Findings)
	assert.String(t, "no fix", d.Findings[0].Fix)
	assert.Integer(t, 1, d.Summary.NoFix)
	assert.False(t, d.Summary.Clean)
}

func TestNeverAffectedReleaseHasNoFinding(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("whiskey", "bookworm", "5.4.1-1", "")
	d, e := o.System().Detail("whiskey")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, d.Findings)
	assert.True(t, d.Summary.Clean)
}

func TestTrackedUnsupportedReleaseIsNeverClean(t *testing.T) {
	o := security_tester.New(t)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("xray", "bullseye", "5.2.5-2.1~deb11u1", "")
	d, e := o.System().Detail("xray")
	assert.FatalOnError(t, e)
	assert.String(t, "unsupported, tracked", d.Summary.Coverage)
	assert.True(t, d.Summary.Checked)
	assert.Count(t, 0, d.Findings)
	assert.False(t, d.Summary.Clean)
}

func TestFixOnTrackedReleaseIsPaidOnly(t *testing.T) {
	o := security_tester.New(t)
	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.KnownIdentifier,
		"xz-utils",
		"bullseye",
		vulnerabilityConstant.DebianResolved,
		"5.2.5-2.1~deb11u9",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("yankee", "bullseye", "5.2.5-2.1~deb11u1", "")
	d, e := o.System().Detail("yankee")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, d.Findings)
	assert.String(t, "paid support only", d.Findings[0].Fix)
	assert.False(t, d.Findings[0].Upgradable)
	assert.Integer(t, 1, d.Summary.PaidOnly)
}

func TestFixInCandidateOnTrackedReleaseIsNotPaidOnly(t *testing.T) {
	o := security_tester.New(t)
	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.KnownIdentifier,
		"xz-utils",
		"bullseye",
		vulnerabilityConstant.DebianResolved,
		"5.2.5-2.1~deb11u9",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("yankee", "bullseye", "5.2.5-2.1~deb11u1", "5.2.5-2.1~deb11u9")
	d, e := o.System().Detail("yankee")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, d.Findings)
	assert.String(t, "debian", d.Findings[0].Fix)
	assert.True(t, d.Findings[0].Upgradable)
	assert.Integer(t, 0, d.Summary.PaidOnly)
	assert.Integer(t, 1, d.Summary.Upgradable)
}

func TestBoardCountsCoverageAndFilters(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("tango", "trixie", "5.6.1-1", "")
	o.PutXz("uniform", "trixie", "5.4.0-1", "5.6.1-1")
	o.PutXz("xray", "bullseye", "5.2.5-2.1~deb11u1", "")
	o.PutXz("zulu", "wheezy", "5.1.1alpha+20120614-2", "")
	b, e := o.System().Board("")
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, b.Total)
	assert.Integer(t, 2, b.Supported)
	assert.Integer(t, 1, b.Tracked)
	assert.Integer(t, 1, b.Untracked)
	assert.Integer(t, 1, b.Upgradable)
	assert.Integer(t, 0, b.Stale)
	assert.Count(t, 4, b.Rows)
	assert.String(t, "uniform", b.Rows[0].Name)
	tracked, f := o.System().Board("tracked")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, tracked.Rows)
	assert.String(t, "xray", tracked.Rows[0].Name)
	assert.Integer(t, 4, tracked.Total)
	_, g := o.System().Board("bogus")
	assert.True(t, validation.Is(g))
}

func TestPackagesGroupFindingsBySourceAndFix(t *testing.T) {
	o := security_tester.New(t)

	for _, v := range [][4]string{
		{
			vulnerabilityConstant.KnownIdentifier,
			"xz-utils",
			vulnerabilityConstant.DebianResolved,
			"5.6.1-1",
		},
		{
			"CVE-2099-10001",
			"xz-utils",
			vulnerabilityConstant.DebianResolved,
			"5.6.2-1",
		},
		{"CVE-2099-10002", "romeo", vulnerabilityConstant.DebianOpen, ""},
		{
			"CVE-2099-10003",
			"sierra",
			vulnerabilityConstant.DebianResolved,
			"2.0-1",
		},
	} {
		o.SecurityTracker.SeedStatus(v[0], v[1], "trixie", v[2], v[3])
	}

	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.FatalOnError(
		t,
		o.System().Put(
			system.New("victor", "trixie", 0),
			[]installed.Installed{
				*installed.New(
					"liblzma5",
					"5.4.0-1",
					"xz-utils",
					"5.4.0-1",
					"5.6.2-1",
				),
				*installed.New("romeo", "1.0-1", "romeo", "1.0-1", ""),
				*installed.New("sierra", "1.0-1", "sierra", "1.0-1", ""),
			},
		),
	)
	p, e := o.System().Packages("victor")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, p.Upgrade)
	assert.String(t, "xz-utils", p.Upgrade[0].Source)
	assert.Count(t, 2, p.Upgrade[0].Vulnerabilities)
	assert.String(t, "5.6.2-1", p.Upgrade[0].Fixed)
	assert.String(t, "5.6.2-1", p.Upgrade[0].Candidate)
	assert.Count(t, 1, p.Pending)
	assert.String(t, "sierra", p.Pending[0].Source)
	assert.Count(t, 1, p.Open)
	assert.String(t, "romeo", p.Open[0].Source)
	assert.Count(t, 0, p.Paid)
	_, f := o.System().Packages("whiskey")
	assert.True(t, not_found.Is(f))
}

func TestPackagesPutPaidOnlyFixUnderPaid(t *testing.T) {
	o := security_tester.New(t)
	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.KnownIdentifier,
		"xz-utils",
		"bullseye",
		vulnerabilityConstant.DebianResolved,
		"5.2.5-2.1~deb11u9",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("xray", "bullseye", "5.2.5-2.1~deb11u1", "")
	p, e := o.System().Packages("xray")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, p.Paid)
	assert.Count(t, 0, p.Open)
	assert.Count(t, 0, p.Upgrade)
}

func TestUntrackedReleaseIsNeverClean(t *testing.T) {
	o := security_tester.New(t)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("zulu", "wheezy", "5.1.1alpha+20120614-2", "")
	d, e := o.System().Detail("zulu")
	assert.FatalOnError(t, e)
	assert.String(t, "unsupported, untracked", d.Summary.Coverage)
	assert.False(t, d.Summary.Checked)
	assert.False(t, d.Summary.Clean)
}

func TestSupportedReleaseIsNotCleanBeforeTrackerLoads(t *testing.T) {
	o := security_tester.New(t)
	o.PutXz("alfa", "trixie", "5.6.1-1", "")
	d, e := o.System().Detail("alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "supported", d.Summary.Coverage)
	assert.False(t, d.Summary.Checked)
	assert.False(t, d.Summary.Clean)
}

func TestSnapshotReplacesPackages(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("bravo", "trixie", "5.4.0-1", "")
	o.PutXz("bravo", "trixie", "5.6.1-1", "")
	d, e := o.System().Detail("bravo")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, d.Findings)
	summaries, f := o.System().Summaries()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, summaries)
}

func TestMalformedSnapshotIsValidation(t *testing.T) {
	s := security_tester.New(t).System()
	assert.True(
		t,
		validation.Is(s.Put(system.New("../charlie", "trixie", 0), nil)),
	)
	assert.True(t, validation.Is(s.Put(system.New("charlie", "", 0), nil)))
}

func TestEmptySnapshotOnTrackedReleaseIsRefused(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	e := o.System().Put(system.New("echo", "trixie", 0), nil)
	assert.True(t, validation.Is(e))
	_, f := o.System().Detail("echo")
	assert.True(t, not_found.Is(f))
	assert.FatalOnError(t, o.System().Put(system.New("echo", "wheezy", 0), nil))
}

func TestPackageWithoutSourceIsRefused(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))

	for _, p := range []*installed.Installed{
		installed.New("liblzma5", "5.4.0-1", "", "5.4.0-1", ""),
		installed.New("liblzma5", "5.4.0-1", "xz-utils", "", ""),
	} {
		e := o.System().Put(
			system.New("foxtrot", "trixie", 0),
			[]installed.Installed{*p},
		)
		assert.True(t, validation.Is(e))
	}

	_, f := o.System().Detail("foxtrot")
	assert.True(t, not_found.Is(f))
}

func TestDeletedSystemIsGone(t *testing.T) {
	o := security_tester.New(t)
	o.PutXz("delta", "trixie", "5.6.1-1", "")
	assert.FatalOnError(t, o.System().Delete("delta"))
	_, e := o.System().Detail("delta")
	assert.True(t, not_found.Is(e))
	assert.True(t, not_found.Is(o.System().Delete("delta")))
}

func TestGroupsCountSystemsPerVulnerability(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.PutXz("echo", "trixie", "5.4.0-1", "5.6.1-1")
	o.PutXz("foxtrot", "trixie", "5.4.0-1", "")
	o.PutXz("golf", "trixie", "5.6.1-1", "")
	groups, e := o.System().Groups()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, groups)
	assert.String(t, "CVE-2024-3094", groups[0].Vulnerability)
	assert.Strings(t, []string{"echo", "foxtrot"}, groups[0].Systems)
	assert.Integer(t, 1, groups[0].Upgradable)
}

func TestGroupsCountSystemOncePerVulnerability(t *testing.T) {
	o := security_tester.New(t)

	for _, source := range []string{"xz-utils", "romeo"} {
		o.SecurityTracker.SeedStatus(
			vulnerabilityConstant.KnownIdentifier,
			source,
			"trixie",
			vulnerabilityConstant.DebianOpen,
			"",
		)
	}

	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.FatalOnError(
		t,
		o.System().Put(
			system.New("sierra", "trixie", 0),
			[]installed.Installed{
				*installed.New("liblzma5", "5.8.1-1", "xz-utils", "5.8.1-1", ""),
				*installed.New("romeo", "1.0-1", "romeo", "1.0-1", ""),
			},
		),
	)
	d, e := o.System().Detail("sierra")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, d.Findings)
	groups, f := o.System().Groups()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, groups)
	assert.Strings(t, []string{"sierra"}, groups[0].Systems)
	assert.Integer(t, 1, groups[0].NoFix)
}

func TestSummariesSortByPriorityThenFindings(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	s := o.System()
	fixed := []installed.Installed{
		*installed.New("liblzma5", "5.6.1-1", "xz-utils", "5.6.1-1", ""),
	}
	assert.FatalOnError(t, s.Put(system.New("hotel", "trixie", 2), fixed))
	assert.FatalOnError(t, s.Put(system.New("india", "trixie", 1), fixed))
	o.PutXz("juliett", "trixie", "5.4.0-1", "")
	summaries, e := s.Summaries()
	assert.FatalOnError(t, e)
	assert.String(t, "juliett", summaries[0].Name)
	assert.String(t, "india", summaries[1].Name)
	assert.String(t, "hotel", summaries[2].Name)
}
