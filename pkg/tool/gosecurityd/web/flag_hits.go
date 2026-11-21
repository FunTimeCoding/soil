package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
)

func (s *Server) flagHits(r *overview_result.Row) (*affected_count.Count, error) {
	hits, e := s.system.Hits([]string{r.Flag.Subject})

	if e != nil {
		return nil, e
	}

	return hits[r.Flag.Subject], nil
}
