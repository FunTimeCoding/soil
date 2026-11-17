package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func (s *Server) Store() *store.Store {
	s.t.Helper()

	if s.store != nil {
		return s.store
	}

	s.Daemon()
	m := lite.New(logger.New(context.Background()), s.litePath)
	s.t.Cleanup(func() { relational.CloseMapper(m) })
	s.store = store.New(m)

	return s.store
}
