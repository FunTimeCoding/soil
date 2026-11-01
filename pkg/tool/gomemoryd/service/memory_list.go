package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
)

func (s *Service) MemoryList(
	memoryType string,
	tag string,
	scope string,
	page int,
) (*memory_list.List, error) {
	result := memory_list.New(
		memoryType,
		tag,
		scope,
		max(page, 1),
		constant.MemoryPageLimit,
	)
	memories, e := s.store.ListMemories(memoryType, tag, listScope(scope), true)

	if e != nil {
		return nil, e
	}

	visible := make([]record.MemorySummary, 0, len(memories))

	for _, m := range memories {
		if !s.isHidden(m.Tags) {
			visible = append(visible, m)
		}
	}

	result.Total = len(visible)
	offset := (result.Page - 1) * result.Limit

	if offset < result.Total {
		result.Memories = visible[offset:min(offset+result.Limit, result.Total)]
	}

	result.Scopes, e = s.store.ListScopes()

	if e != nil {
		return nil, e
	}

	return result, nil
}
