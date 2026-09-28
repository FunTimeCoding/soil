package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"testing"
	"time"
)

func TestSweepRemovesOnlyStalePlacements(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.Save(
		t,
		s,
		constant.SourceKubernetes,
		constant.FixtureService,
		now,
	)
	store_tester.Save(
		t,
		s,
		constant.SourceKubernetes,
		"retired",
		now.Add(-8*24*time.Hour),
	)
	swept, e := s.SweepPlacements(
		constant.SourceKubernetes,
		now.Add(-7*24*time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(1), swept)
	remaining, f := s.Placements()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
	assert.String(t, "relay", remaining[0].Name)
}

func TestSweepLeavesOtherSources(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.Save(
		t,
		s,
		constant.SourceProcess,
		"gosublimed",
		now.Add(-8*24*time.Hour),
	)
	swept, e := s.SweepPlacements(
		constant.SourceKubernetes,
		now.Add(-7*24*time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(0), swept)
	remaining, f := s.Placements()
	assert.FatalOnError(t, f)
	assert.Count(t, 1, remaining)
}

func TestSavePlacementUpsertsOnAttribution(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.Save(
		t,
		s,
		constant.SourceKubernetes,
		constant.FixtureService,
		now,
	)
	store_tester.Save(
		t,
		s,
		constant.SourceKubernetes,
		constant.FixtureService,
		now.Add(time.Minute),
	)
	remaining, e := s.Placements()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, remaining)
}
