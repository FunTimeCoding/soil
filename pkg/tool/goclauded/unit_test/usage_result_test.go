package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/usage_result"
	"testing"
	"time"
)

func result(fiveHourReset time.Time) *usage_result.Result {
	return usage_result.New(
		26,
		fiveHourReset,
		20,
		time.Date(2026, 9, 2, 21, 0, 0, 0, time.UTC),
		34,
		"Wed 8:59 PM",
		time.Now(),
	)
}

func TestFiveHourResetTextMinutes(t *testing.T) {
	assert.String(
		t,
		"in 33 min",
		result(time.Now().Add(33*time.Minute+time.Second)).FiveHourResetText(),
	)
}

func TestFiveHourResetTextHours(t *testing.T) {
	assert.String(
		t,
		"in 2 hr 5 min",
		result(
			time.Now().Add(2*time.Hour+5*time.Minute+time.Second),
		).FiveHourResetText(),
	)
}

func TestFiveHourResetTextPassed(t *testing.T) {
	assert.String(
		t,
		"now",
		result(time.Now().Add(-time.Minute)).FiveHourResetText(),
	)
}

func TestSevenDayResetText(t *testing.T) {
	assert.String(t, "Wed 21:00", result(time.Now()).SevenDayResetText())
}

func TestHasFable(t *testing.T) {
	assert.True(t, result(time.Now()).HasFable())
}
