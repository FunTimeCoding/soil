package conversation_page

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func relativeTimestamp(timestamp string) string {
	t, e := time.Parse(time.RFC3339Nano, timestamp)

	if e != nil {
		return timestamp
	}

	return soilTime.Relative(t)
}
