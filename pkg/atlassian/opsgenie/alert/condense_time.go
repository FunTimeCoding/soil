package alert

import (
	"fmt"
	"github.com/docker/go-units"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func condenseTime(t time.Time) string {
	var format string
	local := t.Local()

	if time.Since(t) < 24*time.Hour {
		format = local.Format(library.HourMinute)
	} else {
		format = local.Format(library.DateMinute)
	}

	return fmt.Sprintf(
		"%s (%s ago)",
		format,
		units.HumanDuration(time.Since(t)),
	)
}
