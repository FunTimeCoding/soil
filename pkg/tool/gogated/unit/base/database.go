package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"gorm.io/gorm"
)

func (s *Server) Database() *gorm.DB {
	s.t.Helper()

	if s.database != nil {
		return s.database
	}

	s.Daemon()
	s.database = lite.New(logger.New(context.Background()), s.litePath)
	s.t.Cleanup(func() { relational.CloseMapper(s.database) })

	return s.database
}
