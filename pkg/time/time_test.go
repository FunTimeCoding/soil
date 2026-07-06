package time

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	assert.Integer(t, 3600, HourInSeconds)
	assert.Integer(t, 86400, DayInSeconds)
	assert.Integer(t, 604800, WeekInSeconds)
	assert.Integer(t, 2419200, MonthInSeconds)
	assert.Integer(t, 29030400, YearInSeconds)
	assert.Integer(t, 28, MonthInDays)
	assert.String(
		t,
		"1970-01-01 00:00:00",
		constant.StartOfTime.Format(DateSecond),
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
