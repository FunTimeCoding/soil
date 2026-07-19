package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	assert.Integer(t, 3600, library.HourInSeconds)
	assert.Integer(t, 86400, library.DayInSeconds)
	assert.Integer(t, 604800, library.WeekInSeconds)
	assert.Integer(t, 2419200, library.MonthInSeconds)
	assert.Integer(t, 29030400, library.YearInSeconds)
	assert.Integer(t, 28, library.MonthInDays)
	assert.String(
		t,
		"1970-01-01 00:00:00",
		constant.StartOfTime.Format(library.DateSecond),
	)
	now := time.Now()
	past := now.Add(-time.Minute)

	if !past.Before(now) {
		t.Fatalf("Past not before present")
	}

	future := now.Add(time.Minute)

	if !future.After(now) {
		t.Fatalf("Future not after present")
	}
}
