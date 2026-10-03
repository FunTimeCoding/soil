package holiday

import (
	"github.com/funtimecoding/soil/pkg/time/day"
	"time"
)

// Reference: https://en.wikipedia.org/wiki/Date_of_Easter#Anonymous_Gregorian_algorithm
func Easter(year int) time.Time {
	golden := year % 19
	century := year / 100
	remainder := year % 100
	leap := century / 4
	leapRemainder := century % 4
	correction := (century + 8) / 25
	moon := (century - correction + 1) / 3
	epact := (19*golden + century - leap - moon + 15) % 30
	quarter := remainder / 4
	quarterRemainder := remainder % 4
	weekday := (32 + 2*leapRemainder + 2*quarter - epact - quarterRemainder) % 7
	shift := (golden + 11*epact + 22*weekday) / 451
	total := epact + weekday - 7*shift + 114

	return day.New(year, time.Month(total/31), total%31+1)
}
