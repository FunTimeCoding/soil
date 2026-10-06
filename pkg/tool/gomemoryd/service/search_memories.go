package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) SearchMemories(
	query string,
	limit int,
	memoryType string,
	tag string,
	scope string,
) ([]record.SearchResult, error) {
	return s.store.SearchMemories(query, limit, memoryType, tag, scope)
}
