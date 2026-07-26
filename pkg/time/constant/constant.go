package constant

const (
	HourMinute       = "15:04"
	HourMinuteSecond = "15:04:05"
	DateYear         = "2006-01-02"
	DateMinute       = "2006-01-02 15:04"
	DateSecond       = "2006-01-02 15:04:05"
	Micro            = "15:04:05.000000"

	HourInMinutes  int = 60
	DayInMinutes       = HourInMinutes * 24
	WeekInMinutes      = DayInMinutes * 7
	MonthInMinutes     = WeekInMinutes * 4
	YearInMinutes      = MonthInMinutes * 12

	MinuteInSeconds int = 60
	HourInSeconds       = HourInMinutes * MinuteInSeconds
	DayInSeconds        = DayInMinutes * MinuteInSeconds
	WeekInSeconds       = WeekInMinutes * MinuteInSeconds
	MonthInSeconds      = MonthInMinutes * MinuteInSeconds
	YearInSeconds       = YearInMinutes * MinuteInSeconds

	MonthInWeeks int = 4

	WeekInDays  int = 7
	MonthInDays     = WeekInDays * MonthInWeeks
)
