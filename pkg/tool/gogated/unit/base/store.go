package base

import "github.com/funtimecoding/soil/pkg/tool/gogated/store"

func (s *Server) Store() *store.Store {
	s.t.Helper()
	s.Service()

	return s.store
}
