package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"testing"
	"time"
)

func TestRecentAlertWindowRunsOnServiceClock(t *testing.T) {
	s := base.New(t, service_tester.WithClock(clock_tester.Window))
	d := relational.Open(logger.New(t.Context()), "", s.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	assert.FatalOnError(
		t,
		store.New(d).Create(
			record.Record{
				Fingerprint: "clocked",
				Name:        "HypervisorUnreachable",
				Severity:    "warning",
				Start:       clock_tester.Window().Add(-30 * time.Minute),
			},
		),
	)
	recent, e := s.AlertClient().GetRecentAlerts(
		t.Context(),
		&client.GetRecentAlertsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, recent)
	assert.String(t, "clocked", recent[0].Fingerprint)
}

func TestTopAlertWindowRunsOnServiceClock(t *testing.T) {
	s := base.New(t, service_tester.WithClock(clock_tester.Window))
	d := relational.Open(logger.New(t.Context()), "", s.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	assert.FatalOnError(
		t,
		store.New(d).Create(
			record.Record{
				Fingerprint: "clocked",
				Name:        "HypervisorUnreachable",
				Severity:    "warning",
				Start:       clock_tester.Window().Add(-48 * time.Hour),
			},
		),
	)
	top, e := s.AlertClient().GetTopAlerts(
		t.Context(),
		&client.GetTopAlertsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, top)
	assert.String(t, "HypervisorUnreachable", top[0].Name)
}

func TestRecentAlertWindowGivenOnlyEndReachesBackFromEnd(t *testing.T) {
	s := base.New(t, service_tester.WithClock(clock_tester.Window))
	d := relational.Open(logger.New(t.Context()), "", s.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	assert.FatalOnError(
		t,
		store.New(d).Create(
			record.Record{
				Fingerprint: "earlier",
				Name:        "HypervisorUnreachable",
				Severity:    "warning",
				Start:       clock_tester.Window().Add(-150 * time.Minute),
			},
		),
	)
	end := clock_tester.Window().Add(-2 * time.Hour)
	recent, e := s.AlertClient().GetRecentAlerts(
		t.Context(),
		&client.GetRecentAlertsParams{End: &end},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, recent)
	assert.String(t, "earlier", recent[0].Fingerprint)
}

func TestTopAlertWindowGivenOnlyEndReachesBackFromEnd(t *testing.T) {
	s := base.New(t, service_tester.WithClock(clock_tester.Window))
	d := relational.Open(logger.New(t.Context()), "", s.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	assert.FatalOnError(
		t,
		store.New(d).Create(
			record.Record{
				Fingerprint: "earlier",
				Name:        "HypervisorUnreachable",
				Severity:    "warning",
				Start:       clock_tester.Window().Add(-180 * time.Hour),
			},
		),
	)
	end := clock_tester.Window().Add(-24 * time.Hour)
	top, e := s.AlertClient().GetTopAlerts(
		t.Context(),
		&client.GetTopAlertsParams{End: &end},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, top)
	assert.String(t, "HypervisorUnreachable", top[0].Name)
}
