package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time"
)

func (s *Server) summaryItems() []string {
	last := "never"

	if v := s.worker.LastPoll(); !v.IsZero() {
		last = time.FormatCompact(v)
	}

	status, e := s.service.Status()
	errors.PanicOnError(e)

	return []string{
		fmt.Sprintf("%d records", status.Total),
		fmt.Sprintf("%d firing", status.Firing),
		fmt.Sprintf("polled %s", last),
	}
}
