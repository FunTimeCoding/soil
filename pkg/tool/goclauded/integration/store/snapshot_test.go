package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
	"time"
)

func TestSnapshotTablesIsolated(t *testing.T) {
	s := store_tester.New(t)
	s.Store.SaveRateSnapshot(26, 20, time.Now(), time.Now())
	s.Store.SaveFableSnapshot(34, "Wed 8:59 PM")
	rate, e := s.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.NotNil(t, rate)
	assert.Integer(t, 26, rate.FiveHourPercent)
	assert.Integer(t, 20, rate.SevenDayPercent)
	fable, f := s.Store.LatestFableSnapshot()
	assert.FatalOnError(t, f)
	assert.NotNil(t, fable)
	assert.Integer(t, 34, fable.Percent)
	assert.String(t, "Wed 8:59 PM", fable.Reset)
}

func TestLatestRateSurvivesFableWrite(t *testing.T) {
	s := store_tester.New(t)
	s.Store.SaveRateSnapshot(26, 20, time.Now(), time.Now())
	s.Store.SaveFableSnapshot(34, "Wed 8:59 PM")
	s.Store.SaveFableSnapshot(35, "Wed 8:59 PM")
	rate, e := s.Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.NotNil(t, rate)
	assert.Integer(t, 26, rate.FiveHourPercent)
}
