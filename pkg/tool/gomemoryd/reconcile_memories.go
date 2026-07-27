package gomemoryd

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
)

func reconcileMemories(s *service.Service) {
	memories, e := s.ListMemories("", "", constant.AllScope, true)
	errors.PanicOnError(e)

	for _, m := range memories {
		full, f := s.GetMemory(m.Identifier)
		errors.PanicOnError(f)
		s.MustReindexMemory(full)
	}
}
