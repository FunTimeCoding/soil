package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/time_of_day"
	"testing"
	"time"
	_ "time/tzdata"
)

func TestTimeOfDayLaterTodayIsToday(t *testing.T) {
	now := time.Date(2026, time.May, 4, 7, 30, 0, 0, time.UTC)
	assert.Any(
		t,
		time.Date(2026, time.May, 4, 9, 0, 0, 0, time.UTC),
		time_of_day.New(9, 0, time.UTC).Next(now),
	)
}

func TestTimeOfDayAlreadyPastIsTomorrow(t *testing.T) {
	now := time.Date(2026, time.May, 4, 9, 0, 0, 0, time.UTC)
	assert.Any(
		t,
		time.Date(2026, time.May, 5, 9, 0, 0, 0, time.UTC),
		time_of_day.New(9, 0, time.UTC).Next(now),
	)
}

func TestTimeOfDayKeepsTheWallClockAcrossDaylightSaving(t *testing.T) {
	chicago, e := time.LoadLocation("America/Chicago")
	assert.FatalOnError(t, e)
	now := time.Date(2026, time.March, 7, 10, 0, 0, 0, chicago)
	next := time_of_day.New(9, 0, chicago).Next(now)
	assert.String(t, "2026-03-08T09:00:00-05:00", next.Format(time.RFC3339))
	assert.Any(t, 22*time.Hour, next.Sub(now))
}

func TestTimeOfDayReadsNowInItsOwnLocation(t *testing.T) {
	chicago, e := time.LoadLocation("America/Chicago")
	assert.FatalOnError(t, e)
	now := time.Date(2026, time.May, 5, 3, 30, 0, 0, time.UTC)
	assert.String(
		t,
		"2026-05-04T23:00:00-05:00",
		time_of_day.New(23, 0, chicago).Next(now).Format(time.RFC3339),
	)
}
