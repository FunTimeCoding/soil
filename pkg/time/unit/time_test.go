package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	library "github.com/funtimecoding/soil/pkg/time"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/time/day"
	"github.com/funtimecoding/soil/pkg/time/minute"
	"github.com/funtimecoding/soil/pkg/time/second"
	"testing"
	"time"
)

func TestConstant(t *testing.T) {
	assert.String(t, "15:04", timeConstant.HourMinute)
	assert.String(t, "15:04:05", timeConstant.HourMinuteSecond)
}

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

func TestFormat(t *testing.T) {
	assert.String(t, "1970-01-01 00:00", library.Format(constant.StartOfTime))
}

func TestFormatCompactToday(t *testing.T) {
	now := time.Now()
	assert.String(t, now.Local().Format("15:04"), library.FormatCompact(now))
}

func TestFormatCompactOlder(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	assert.String(
		t,
		yesterday.Local().Format("2006-01-02 15:04"),
		library.FormatCompact(yesterday),
	)
}

func TestParse(t *testing.T) {
	result := library.Parse(timeConstant.DateYear, "2019-01-01")
	assert.Integer(t, 2019, result.Year())
	assert.Integer(t, 1, int(result.Month()))
	assert.Integer(t, 1, result.Day())
}

func TestMidnight(t *testing.T) {
	assert.Any(
		t,
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local),
		library.Midnight(constant.StartOfTime),
	)
}

func TestOlderThan(t *testing.T) {
	now := time.Now()
	assert.False(t, library.OlderThan(now, 1))
	fiveMinutesAgo := now.Add(-5 * time.Minute)
	assert.True(t, library.OlderThan(fiveMinutesAgo, 1))
}

func TestRelative(t *testing.T) {
	now := time.Now()
	assert.String(t, "just now", library.Relative(now.Add(-time.Second)))
	assert.String(t, "5m ago", library.Relative(now.Add(-5*time.Minute)))
	assert.String(t, "3h ago", library.Relative(now.Add(-3*time.Hour)))
	assert.String(t, "2d ago", library.Relative(now.Add(-50*time.Hour)))
}

func TestPublicHoliday(t *testing.T) {
	assert.True(t, library.PublicHoliday(day.New(2024, time.November, 1)))
	assert.False(t, library.PublicHoliday(day.New(2024, time.November, 2)))
}

func TestWorkDaysSince(t *testing.T) {
	assert.Integer(
		t,
		3,
		library.WeekDaysSince(
			day.New(2024, time.October, 25),
			day.New(2024, time.October, 30),
		),
	)
}

func TestMinuteReadable(t *testing.T) {
	assert.String(t, "1 minute", minute.Readable(1))
	assert.String(t, "1 hour", minute.Readable(60))
}

func TestSecondReadable(t *testing.T) {
	minuteSeconds := timeConstant.MinuteInSeconds
	hourSeconds := timeConstant.HourInSeconds
	daySeconds := timeConstant.DayInSeconds
	monthSeconds := timeConstant.MonthInSeconds
	assert.String(t, "1 second", second.Readable(1))
	assert.String(t, "59 seconds", second.Readable(59))
	assert.String(t, "1 minute", second.Readable(minuteSeconds))
	assert.String(t, "30 minutes", second.Readable(minuteSeconds*30))
	assert.String(t, "59 minutes", second.Readable(minuteSeconds*59))
	assert.String(t, "1 hour", second.Readable(hourSeconds))
	assert.String(t, "1 day", second.Readable(daySeconds))
	assert.String(t, "1.5 days", second.Readable(hourSeconds*36))
	assert.String(t, "2 days", second.Readable(daySeconds*2))
	assert.String(t, "1 week", second.Readable(daySeconds*7))
	assert.String(t, "1 month", second.Readable(monthSeconds))
}

func TestOfDay(t *testing.T) {
	now := time.Now()
	y, m, d := now.Date()
	midnight := time.Date(y, m, d, 0, 0, 0, 0, now.Location())
	one := time.Date(y, m, d, 1, 0, 0, 0, now.Location())
	assert.Integer(t, 0, second.OfDay(midnight))
	assert.Integer(t, 3600, second.OfDay(one))
}

func TestToClock(t *testing.T) {
	assert.String(t, "00:00", second.ToClock(0))
}
