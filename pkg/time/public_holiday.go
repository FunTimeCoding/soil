package time

import "time"

var PublicHolidays = []time.Time{
	NewDay(2024, time.November, 1),  // All Saints' Day
	NewDay(2024, time.December, 25), // Christmas Day
	NewDay(2024, time.December, 26), // St. Stephen's Day
	NewDay(2025, time.January, 1),   // New Year's Day
	NewDay(2025, time.January, 6),   // Epiphany
	NewDay(2025, time.April, 18),    // Good Friday
	NewDay(2025, time.April, 21),    // Easter Monday
	NewDay(2025, time.May, 1),       // Labour Day
	NewDay(2025, time.May, 29),      // Ascension Day
	NewDay(2025, time.June, 9),      // Whit Monday
	NewDay(2025, time.June, 19),     // Corpus Christi
	NewDay(2025, time.October, 3),   // German Unity Day
	NewDay(2025, time.November, 1),  // All Saints' Day
	NewDay(2025, time.December, 25), // Christmas Day
	NewDay(2025, time.December, 26), // St. Stephen's Day
}

func PublicHoliday(t time.Time) bool {
	for _, d := range PublicHolidays {
		if d.Year() == t.Year() &&
			d.Month() == t.Month() &&
			d.Day() == t.Day() {
			return true
		}
	}

	return false
}
