package issue

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func (i *Issue) AgeWeekDays() int {
	return soilTime.WeekDaysSince(i.ChangeTime(), time.Now())
}
