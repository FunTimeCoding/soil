package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
)

func (s *Server) summaryItems() []string {
	last := "never"

	if v := s.worker.LastPoll(); !v.IsZero() {
		last = time.FormatCompact(v)
	}

	return []string{
		fmt.Sprintf("%d records", s.store.MustCount()),
		fmt.Sprintf("%d firing", s.worker.FiringCount()),
		fmt.Sprintf("polled %s", last),
	}
}
