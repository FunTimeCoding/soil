package dashboard_page

import (
	"fmt"
	"time"
)

func duration(d time.Duration) string {
	if d == 0 {
		return "-"
	}

	if d < time.Minute {
		return d.Round(time.Second).String()
	}

	if d < time.Hour {
		return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
	}

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60

	return fmt.Sprintf("%dh%dm", hours, minutes)
}
