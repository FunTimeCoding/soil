package timeline

import (
	"fmt"
	"time"
)

func Format(e *Entry) string {
	t, f := time.Parse(time.RFC3339, e.Timestamp)
	timestamp := e.Timestamp

	if f == nil {
		timestamp = t.Local().Format("Jan 02 15:04")
	}

	if e.Identifier == 0 {
		return fmt.Sprintf("%s  %s", timestamp, FormatDescription(e))
	}

	return fmt.Sprintf(
		"%s  [%d] %s",
		timestamp,
		e.Identifier,
		FormatDescription(e),
	)
}
