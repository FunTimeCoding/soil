package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_view"
)

func (s *Service) MemoryView(identifier int64) (*memory_view.View, error) {
	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if s.isHidden(m.Tags) {
		return nil, not_found.New(constant.MemoryKind, identifier)
	}

	result := memory_view.New(m)
	result.History, e = s.store.GetMemoryHistory(identifier)

	if e != nil {
		return nil, e
	}

	related, f := s.store.ListRelated(identifier)

	if f != nil {
		return nil, f
	}

	for _, r := range related {
		if !s.isHidden(r.Tags) {
			result.Related = append(result.Related, r)
		}
	}

	return result, nil
}
