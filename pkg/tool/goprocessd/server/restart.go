package server

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (s *Server) Restart(names []string) error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()

	for _, name := range names {
		p := s.findProcess(name)

		if p == nil {
			return validation.New(constant.UnknownProcess, name)
		}

		if e := p.Stop(); e != nil {
			return e
		}

		s.spawn(p)
	}

	return nil
}
