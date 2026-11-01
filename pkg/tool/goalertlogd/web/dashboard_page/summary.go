package dashboard_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/status"
	stdTime "time"
)

func Summary(
	s *status.Status,
	lastPoll stdTime.Time,
) []string {
	last := "never"

	if !lastPoll.IsZero() {
		last = time.FormatCompact(lastPoll)
	}

	return []string{
		fmt.Sprintf("%d records", s.Total),
		fmt.Sprintf("%d firing", s.Firing),
		fmt.Sprintf("polled %s", last),
	}
}
