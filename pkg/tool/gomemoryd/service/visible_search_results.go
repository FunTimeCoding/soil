package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) VisibleSearchResults(
	query string,
	limit int,
	memoryType string,
	tag string,
	scope string,
) ([]record.SearchResult, error) {
	results, e := s.store.SearchMemories(query, limit, memoryType, tag, scope)

	if e != nil {
		return nil, e
	}

	result := make([]record.SearchResult, 0, len(results))

	for _, v := range results {
		if s.isHidden(v.Tags) {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
