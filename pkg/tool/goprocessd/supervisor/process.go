package supervisor

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
)

func (s *Supervisor) process(name string) (*process.Process, error) {
	if p := s.findProcess(name); p != nil {
		return p, nil
	}

	return nil, not_found.New(constant.ProcessKind, name)
}
