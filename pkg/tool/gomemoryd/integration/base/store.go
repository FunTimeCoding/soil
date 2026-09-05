package base

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Server) Store() *store.Store {
	return s.store
}
