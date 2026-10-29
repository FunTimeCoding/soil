package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Server) summaryItems() []string {
	count, e := s.service.Count()
	errors.PanicOnError(e)

	return []string{fmt.Sprintf("%d entries", count)}
}
