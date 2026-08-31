package usage_result

import (
	"fmt"
	"time"
)

func (r *Result) FiveHourResetText() string {
	remaining := time.Until(r.FiveHourReset).Round(time.Minute)

	if remaining <= 0 {
		return "now"
	}

	if remaining < time.Hour {
		return fmt.Sprintf("in %d min", int(remaining.Minutes()))
	}

	return fmt.Sprintf(
		"in %d hr %d min",
		int(remaining.Hours()),
		int(remaining.Minutes())%60,
	)
}
