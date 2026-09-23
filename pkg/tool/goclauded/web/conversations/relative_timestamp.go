package conversations

import (
	moment "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func relativeTimestamp(timestamp string) string {
	t, e := time.Parse(time.RFC3339Nano, timestamp)

	if e != nil {
		return timestamp
	}

	return moment.Relative(t)
}
