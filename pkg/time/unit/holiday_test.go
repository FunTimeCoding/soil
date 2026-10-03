package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/day"
	"github.com/funtimecoding/soil/pkg/time/holiday"
	"testing"
	"time"
)

func TestEaster(t *testing.T) {
	assert.Any(t, day.New(2024, time.March, 31), holiday.Easter(2024))
	assert.Any(t, day.New(2025, time.April, 20), holiday.Easter(2025))
	assert.Any(t, day.New(2026, time.April, 5), holiday.Easter(2026))
	assert.Any(t, day.New(2027, time.March, 28), holiday.Easter(2027))
	assert.Any(t, day.New(2038, time.April, 25), holiday.Easter(2038))
	assert.Any(t, day.New(2285, time.March, 22), holiday.Easter(2285))
}

func TestPublicHolidayCalendarFrom2024To2027(t *testing.T) {
	assert.Any(
		t,
		[]time.Time{
			day.New(2024, time.November, 1),
			day.New(2024, time.December, 25),
			day.New(2024, time.December, 26),
			day.New(2025, time.January, 1),
			day.New(2025, time.January, 6),
			day.New(2025, time.April, 18),
			day.New(2025, time.April, 21),
			day.New(2025, time.May, 1),
			day.New(2025, time.May, 29),
			day.New(2025, time.June, 9),
			day.New(2025, time.June, 19),
			day.New(2025, time.October, 3),
			day.New(2025, time.November, 1),
			day.New(2025, time.December, 25),
			day.New(2025, time.December, 26),
			day.New(2026, time.January, 1),
			day.New(2026, time.January, 6),
			day.New(2026, time.April, 3),
			day.New(2026, time.April, 6),
			day.New(2026, time.May, 1),
			day.New(2026, time.May, 14),
			day.New(2026, time.May, 25),
			day.New(2026, time.June, 4),
			day.New(2026, time.October, 3),
			day.New(2026, time.November, 1),
			day.New(2026, time.December, 25),
			day.New(2026, time.December, 26),
			day.New(2027, time.January, 1),
			day.New(2027, time.January, 6),
			day.New(2027, time.March, 26),
			day.New(2027, time.March, 29),
			day.New(2027, time.May, 1),
			day.New(2027, time.May, 6),
			day.New(2027, time.May, 17),
			day.New(2027, time.May, 27),
			day.New(2027, time.October, 3),
			day.New(2027, time.November, 1),
			day.New(2027, time.December, 25),
			day.New(2027, time.December, 26),
		},
		holidaysBetween(
			day.New(2024, time.November, 1),
			day.New(2027, time.December, 31),
		),
	)
}
