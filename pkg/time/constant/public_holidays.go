package constant

import (
	"github.com/funtimecoding/soil/pkg/time/day"
	"time"
)

var PublicHolidays = []time.Time{
	day.New(2024, time.November, 1),  // All Saints' Day
	day.New(2024, time.December, 25), // Christmas Day
	day.New(2024, time.December, 26), // St. Stephen's Day
	day.New(2025, time.January, 1),   // New Year's Day
	day.New(2025, time.January, 6),   // Epiphany
	day.New(2025, time.April, 18),    // Good Friday
	day.New(2025, time.April, 21),    // Easter Monday
	day.New(2025, time.May, 1),       // Labour Day
	day.New(2025, time.May, 29),      // Ascension Day
	day.New(2025, time.June, 9),      // Whit Monday
	day.New(2025, time.June, 19),     // Corpus Christi
	day.New(2025, time.October, 3),   // German Unity Day
	day.New(2025, time.November, 1),  // All Saints' Day
	day.New(2025, time.December, 25), // Christmas Day
	day.New(2025, time.December, 26), // St. Stephen's Day
}
