package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.stop.Close()
			s.memoryGuard.Close()
			s.queryGuard.Close()
			s.closeBrowser()
			s.memory.Close()
			s.query.Close()
			errors.PanicClose(s.searchDatabase)
			relational.CloseMapper(s.database)
		},
	)
}
