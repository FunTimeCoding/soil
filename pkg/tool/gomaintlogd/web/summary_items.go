package web

import "fmt"

func (s *Server) summaryItems() []string {
	return []string{fmt.Sprintf("%d entries", s.store.Count())}
}
