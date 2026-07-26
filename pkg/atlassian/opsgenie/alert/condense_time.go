package alert

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func condenseTime(t time.Time) string {
	var format string
	local := t.Local()

	if time.Since(t) < 24*time.Hour {
		format = local.Format(constant.HourMinute)
	} else {
		format = local.Format(constant.DateMinute)
	}

	return fmt.Sprintf(
		"%s (%s ago)",
		format,
		units.HumanDuration(time.Since(t)),
	)
}
