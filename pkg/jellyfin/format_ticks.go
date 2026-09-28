package jellyfin

import "fmt"

func formatTicks(ticks int64) string {
	seconds := ticks / 10_000_000
	minutes := seconds / 60
	hours := minutes / 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes%60, seconds%60)
	}

	return fmt.Sprintf("%d:%02d", minutes, seconds%60)
}
