package base

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"gorm.io/gorm"
)

func (s *Server) Mapper() *gorm.DB {
	s.t.Helper()

	if s.mapper != nil {
		return s.mapper
	}

	s.Daemon()
	s.mapper = lite.New(logger.New(s.t.Context()), s.litePath)
	s.t.Cleanup(func() { relational.CloseMapper(s.mapper) })

	return s.mapper
}
