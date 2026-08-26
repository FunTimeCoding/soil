package server

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (s *Server) LogLines(
	name string,
	all bool,
) ([]string, int, error) {
	p := s.findProcess(name)

	if p == nil {
		return nil, 0, validation.New(constant.UnknownProcess, name)
	}

	if all {
		return p.Log(), 0, nil
	}

	current, older := p.CurrentLog()

	return current, older, nil
}
