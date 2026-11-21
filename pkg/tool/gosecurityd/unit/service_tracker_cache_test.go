package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRestartAnswersFromCacheWithoutDownload(t *testing.T) {
	first := security_tester.New(t)
	first.SecurityTracker.SeedKnown()
	first.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, first.Debian().Refresh(t.Context()))
	second := security_tester.New(t)
	second.Cache = first.Cache
	second.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, second.Debian().Refresh(t.Context()))
	assert.Integer(t, 0, second.SecurityTracker.Downloads())
	r, e := second.Debian().Status(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, r.Statuses)
}

func TestCachedTrackerCountsAsVerifiedAtItsModifiedTime(t *testing.T) {
	modified := time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC)
	first := security_tester.New(t)
	first.SecurityTracker.SeedKnown()
	first.SecurityTracker.Modify(modified)
	assert.Count(t, 0, first.Debian().Refresh(t.Context()))
	second := security_tester.New(t)
	second.Cache = first.Cache
	second.SecurityTracker.Fail(errors.New("romeo"))
	assert.Count(t, 1, second.Debian().Refresh(t.Context()))
	assert.Time(t, modified, second.Debian().Covering("trixie").Verified())
}

func TestFailedRefreshDoesNotVerifyTracker(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	verified := o.Debian().Covering("trixie").Verified()
	assert.False(t, verified.IsZero())
	o.Later = time.Hour
	o.SecurityTracker.Fail(errors.New("sierra"))
	assert.Count(t, 1, o.Debian().Refresh(t.Context()))
	assert.Time(t, verified, o.Debian().Covering("trixie").Verified())
	o.SecurityTracker.Fail(nil)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.True(t, o.Debian().Covering("trixie").Verified().After(verified))
}

func TestUnchangedTrackerIsVerifiedAgain(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	verified := o.Debian().Covering("trixie").Verified()
	o.Later = time.Hour
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.Integer(t, 1, o.SecurityTracker.Downloads())
	assert.True(t, o.Debian().Covering("trixie").Verified().After(verified))
}

func TestUnchangedTrackerIsNotDownloadedAgain(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.Integer(t, 1, o.SecurityTracker.Downloads())
}

func TestChangedTrackerIsDownloadedAgain(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.Modify(
		time.Date(2026, time.November, 19, 9, 0, 0, 0, time.UTC),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.Integer(t, 2, o.SecurityTracker.Downloads())
	r, e := o.Debian().Status(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, r.Statuses)
}

func TestDamagedCacheIsReplaced(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	assert.FatalOnError(
		t,
		os.WriteFile(
			filepath.Join(o.Cache, constant.DebianCache),
			[]byte("sierra"),
			0o600,
		),
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	assert.Integer(t, 1, o.SecurityTracker.Downloads())
	_, e := o.Debian().Status(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, e)
}

func TestExtendedTrackerAnswersUnsupportedReleases(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.SeedStatus(
		vulnerabilityConstant.FixtureIdentifier,
		"tango-utils",
		"bullseye",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.FixtureIdentifier,
		"tango-utils",
		"bullseye",
		vulnerabilityConstant.DebianResolved,
		"8.2.5-2.1~deb11u9",
	)
	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.FixtureIdentifier,
		"tango-utils",
		"trixie",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	r, e := o.Debian().Status(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 3, r.Statuses)
	bullseye, f := o.Debian().Status(
		vulnerabilityConstant.FixtureIdentifier,
		"bullseye",
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, bullseye.Statuses)
	assert.String(t, "resolved", bullseye.Statuses[0].Status)
	assert.String(t, "8.2.5-2.1~deb11u9", bullseye.Statuses[0].Fixed)
	trixie, g := o.Debian().Status(
		vulnerabilityConstant.FixtureIdentifier,
		"trixie",
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, trixie.Statuses)
	assert.String(t, "resolved", trixie.Statuses[0].Status)
}
