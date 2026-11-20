package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"testing"
	"time"
)

func TestSavePackageRefreshesVersion(t *testing.T) {
	o := store_tester.New(t)
	now := time.Now()
	o.SavePackage("goagentd.service", "goagentd", "0.2.122", now)
	o.SavePackage(
		"goagentd.service",
		"goagentd",
		"0.11.154",
		now.Add(time.Minute),
	)
	v, e := o.Store.Placements()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestMatchingPlacementsFilterByPackage(t *testing.T) {
	o := store_tester.New(t)
	now := time.Now()
	o.SavePackage("goagentd.service", "goagentd", "0.11.154", now)
	o.SavePackage("bravo.service", "bravo", "0.11.154", now)
	v, e := o.Store.MatchingPlacements("", constant.SourceOutpost, "goagentd")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd", v[0].Package)
}
