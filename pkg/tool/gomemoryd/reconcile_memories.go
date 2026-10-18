package gomemoryd

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
)

func reconcileMemories(s *service.Service) []error {
	memories, e := s.ListMemories("", "", constant.AllScope, true)
	errors.PanicOnError(e)
	var failed []error

	for _, m := range memories {
		full, f := s.GetMemory(m.Identifier)
		errors.PanicOnError(f)

		if g := s.ReindexMemory(full); g != nil {
			failed = append(failed, g)

			if connection.Classify(g) != nil {
				return failed
			}
		}
	}

	return failed
}
