package constant

import (
	"github.com/funtimecoding/soil/pkg/time/holiday"
	"time"
)

var (
	NewYearsDay      = holiday.NewFixed(time.January, 1)
	Epiphany         = holiday.NewFixed(time.January, 6)
	GoodFriday       = holiday.NewEaster(-2)
	EasterMonday     = holiday.NewEaster(1)
	LabourDay        = holiday.NewFixed(time.May, 1)
	AscensionDay     = holiday.NewEaster(39)
	WhitMonday       = holiday.NewEaster(50)
	CorpusChristi    = holiday.NewEaster(60)
	GermanUnityDay   = holiday.NewFixed(time.October, 3)
	AllSaintsDay     = holiday.NewFixed(time.November, 1)
	ChristmasDay     = holiday.NewFixed(time.December, 25)
	SaintStephensDay = holiday.NewFixed(time.December, 26)

	PublicHolidays = []*holiday.Holiday{
		NewYearsDay,
		Epiphany,
		GoodFriday,
		EasterMonday,
		LabourDay,
		AscensionDay,
		WhitMonday,
		CorpusChristi,
		GermanUnityDay,
		AllSaintsDay,
		ChristmasDay,
		SaintStephensDay,
	}
)
