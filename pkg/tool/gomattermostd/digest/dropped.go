package digest

import (
	"fmt"
	"strings"
	"time"
)

func Dropped(label []string, window time.Duration) string {
	return fmt.Sprintf(
		"dropped %s after %d days without activity - subscribe again if still relevant",
		strings.Join(label, ", "),
		int(window.Hours()/24),
	)
}
