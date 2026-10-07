package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func (s *Service) ListDocuments(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) (*result.ListOutcome, error) {
	results, e := s.store.ListDocuments(
		collection,
		metadata,
		limit,
		offset,
		full,
	)

	if e != nil {
		return nil, e
	}

	return result.NewListOutcome(
		results,
		s.store.CollectionFacets(collection, metadata, 20),
	), nil
}
