package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) SearchMemories(
	query string,
	limit int,
	memoryType string,
	tag string,
) ([]store.SearchResult, error) {
	return s.store.SearchMemories(query, limit, memoryType, tag)
}
