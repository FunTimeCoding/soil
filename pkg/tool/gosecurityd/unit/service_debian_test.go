package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestStatusBeforeFirstLoadIsUnavailable(t *testing.T) {
	_, e := security_tester.New(t).Debian().Status(
		vulnerabilityConstant.KnownIdentifier,
		"",
	)
	assert.True(t, unavailable.Is(e))
}

func TestRefreshLoadsTracker(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	r, e := s.Status(vulnerabilityConstant.KnownIdentifier, "")
	assert.FatalOnError(t, e)
	assert.String(t, "CVE-2024-3094", r.Identifier)
	assert.Count(t, 2, r.Statuses)
}

func TestReleaseNarrowsStatuses(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	r, e := s.Status(vulnerabilityConstant.KnownIdentifier, "trixie")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.Statuses)
	assert.String(t, "trixie", r.Statuses[0].Release)
}

func TestListedReleaseWithoutEntryAnswersEmpty(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.SeedStatus(
		"CVE-2023-38703",
		"asterisk",
		"sid",
		vulnerabilityConstant.DebianResolved,
		"1:20.8.1-1",
	)
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	r, e := s.Status(vulnerabilityConstant.KnownIdentifier, "sid")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, r.Statuses)
}

func TestUnlistedReleaseIsNotFound(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	_, e := s.Status(vulnerabilityConstant.KnownIdentifier, "bullseye")
	assert.True(t, not_found.Is(e))
	assert.String(
		t,
		"release not found: bullseye (tracker lists bookworm, trixie)",
		e.Error(),
	)
}

func TestUnknownIdentifierInTrackerIsNotFound(t *testing.T) {
	o := security_tester.New(t)
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	_, e := s.Status("CVE-2099-99999", "")
	assert.True(t, not_found.Is(e))
}

func TestRefreshSwapsInNewTracker(t *testing.T) {
	o := security_tester.New(t)
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	o.SecurityTracker.SeedKnown()
	assert.Count(t, 0, s.Refresh(t.Context()))
	r, e := s.Status(vulnerabilityConstant.KnownIdentifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, r.Statuses)
}

func TestFailedRefreshKeepsPreviousTracker(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedKnown()
	s := o.Debian()
	assert.Count(t, 0, s.Refresh(t.Context()))
	o.SecurityTracker.Fail(errors.New("charlie"))
	assert.Count(t, 1, s.Refresh(t.Context()))
	r, e := s.Status(vulnerabilityConstant.KnownIdentifier, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, r.Statuses)
}

func TestMalformedIdentifierIsRefusedByStatus(t *testing.T) {
	_, e := security_tester.New(t).Debian().Status("delta", "")
	assert.True(t, validation.Is(e))
}
