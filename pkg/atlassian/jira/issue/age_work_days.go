package issue

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func (i *Issue) AgeWorkDays() int {
	return soilTime.WorkDays(i.ChangeTime(), time.Now())
}
