package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func (s *Service) Reconcile(x context.Context) []error {
	memories, e := s.store.ListMemories("", "", constant.AllScope, false)

	if e != nil {
		return []error{e}
	}

	var failed []error

	for _, m := range memories {
		full, f := s.store.GetMemory(m.Identifier)

		if f != nil {
			return append(failed, f)
		}

		if g := s.syncIndex(x, full); g != nil {
			failed = append(failed, g)

			if unavailable.Is(g) || x.Err() != nil {
				return failed
			}
		}
	}

	return failed
}
