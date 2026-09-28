package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"testing"
	"time"
)

func TestSavePackageRefreshesTheVersion(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.SavePackage(
		t,
		s,
		"goagentd.service",
		"goagentd",
		"0.2.122",
		now,
	)
	store_tester.SavePackage(
		t,
		s,
		"goagentd.service",
		"goagentd",
		"0.11.154",
		now.Add(time.Minute),
	)
	v, e := s.Placements()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestMatchingPlacementsFilterByPackage(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.SavePackage(
		t,
		s,
		"goagentd.service",
		"goagentd",
		"0.11.154",
		now,
	)
	store_tester.SavePackage(t, s, "bravo.service", "bravo", "0.11.154", now)
	v, e := s.MatchingPlacements("", constant.SourceOutpost, "goagentd")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd", v[0].Package)
}
