package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
	"time"
)

func findingsByKind(
	t *testing.T,
	s *service_tester.Tester,
	kind string,
) []*finding.Finding {
	t.Helper()
	all, e := s.Service.Findings()
	assert.FatalOnError(t, e)
	var result []*finding.Finding

	for _, i := range all {
		if i.Kind == kind {
			result = append(result, i)
		}
	}

	return result
}

func TestFindingsCleanState(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "healthy", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	result, e := s.Service.Findings()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}

func TestFindingsPoolExhausted(t *testing.T) {
	s := service_tester.New(t)
	assert.Count(t, 0, findingsByKind(t, s, constant.PoolExhausted))

	for i := 0; ; i++ {
		r := s.Register(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}
	}

	result := findingsByKind(t, s, constant.PoolExhausted)
	assert.Count(t, 1, result)
	assert.String(t, "", result[0].Subject)
}

func TestFindingsIgnoreLiveSessionAwaitingSweep(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("running")
	assert.Count(t, 0, findingsByKind(t, s, constant.MissingTranscriptEmpty))
	assert.Count(t, 0, findingsByKind(t, s, constant.MissingTranscriptKept))
	s.Store.Advance(31 * time.Minute)
	assert.Count(t, 1, findingsByKind(t, s, constant.MissingTranscriptEmpty))
}

func TestFindingsMissingTranscriptSplit(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	s.Store.EnsureSession("named")
	s.Store.Store.UpdateFields("named", map[string]any{"alias": "kept-by-hand"})
	s.Store.Advance(31 * time.Minute)
	empty := findingsByKind(t, s, constant.MissingTranscriptEmpty)
	assert.Count(t, 1, empty)
	assert.String(t, "husk", empty[0].Subject)
	kept := findingsByKind(t, s, constant.MissingTranscriptKept)
	assert.Count(t, 1, kept)
	assert.String(t, "named", kept[0].Subject)
	assert.String(t, "session has an alias", kept[0].Detail)
}

func TestFindingsOrphanTrackerState(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "vanished", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.Count(t, 0, findingsByKind(t, s, constant.OrphanTrackerState))
	assert.FatalOnError(t, s.Store.Store.DeleteSession("vanished"))
	result := findingsByKind(t, s, constant.OrphanTrackerState)
	assert.Count(t, 1, result)
	assert.String(t, "vanished", result[0].Subject)
}

func TestFindingsUnownedQueue(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"",
			"Nobody",
			constant.QueueTimeout,
			"orphaned",
		),
	)
	result := findingsByKind(t, s, constant.UnownedQueue)
	assert.Count(t, 1, result)
	assert.String(t, "", result[0].Subject)
	assert.Integer(t, 1, result[0].Count)
	assert.StringContains(t, "Nobody", result[0].Detail)
}

func TestFindingsUnownedNotification(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Store.SendNotification("", "Nobody", "Dale", "orphaned"),
	)
	result := findingsByKind(t, s, constant.UnownedNotification)
	assert.Count(t, 1, result)
	assert.String(t, "", result[0].Subject)
	assert.StringContains(t, "Nobody", result[0].Detail)
}

func TestFindingsStaleCallsign(t *testing.T) {
	s := service_tester.New(t)
	r := s.RegisterActive("lingering")
	assert.Count(t, 0, findingsByKind(t, s, constant.StaleCallsign))
	s.Store.Advance(8 * 24 * time.Hour)
	result := findingsByKind(t, s, constant.StaleCallsign)
	assert.Count(t, 1, result)
	assert.String(t, r.Callsign, result[0].Subject)
}

func TestFindingsStaleCallsignClearedBySweep(t *testing.T) {
	s := service_tester.New(t)
	s.RegisterActive("lingering")
	s.Store.Advance(8 * 24 * time.Hour)
	assert.Count(t, 1, findingsByKind(t, s, constant.StaleCallsign))
	s.Service.RunTimeoutSweep()
	assert.Count(t, 0, findingsByKind(t, s, constant.StaleCallsign))
}
