package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	assert.Integer(t, 3600, timeConstant.HourInSeconds)
	assert.Integer(t, 86400, timeConstant.DayInSeconds)
	assert.Integer(t, 604800, timeConstant.WeekInSeconds)
	assert.Integer(t, 2419200, timeConstant.MonthInSeconds)
	assert.Integer(t, 29030400, timeConstant.YearInSeconds)
	assert.Integer(t, 28, timeConstant.MonthInDays)
	assert.String(
		t,
		"1970-01-01 00:00:00",
		constant.StartOfTime.Format(timeConstant.DateSecond),
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
