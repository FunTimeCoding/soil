package constant

const (
	HourMinute       = "15:04"
	HourMinuteSecond = "15:04:05"
	DateYear         = "2006-01-02"
	DateMinute       = "2006-01-02 15:04"
	DateSecond       = "2006-01-02 15:04:05"
	Micro            = "15:04:05.000000"

	HourMinutes  int = 60
	DayMinutes       = HourMinutes * 24
	WeekMinutes      = DayMinutes * 7
	MonthMinutes     = WeekMinutes * 4
	YearMinutes      = MonthMinutes * 12

	MinuteSeconds int = 60
	HourSeconds       = HourMinutes * MinuteSeconds
	DaySeconds        = DayMinutes * MinuteSeconds
	WeekSeconds       = WeekMinutes * MinuteSeconds
	MonthSeconds      = MonthMinutes * MinuteSeconds
	YearSeconds       = YearMinutes * MinuteSeconds

	MonthWeeks int = 4

	WeekDays  int = 7
	MonthDays     = WeekDays * MonthWeeks
)
