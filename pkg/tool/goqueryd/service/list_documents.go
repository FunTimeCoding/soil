package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func (s *Service) ListDocuments(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) (*search.ListOutcome, error) {
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

	return search.NewListOutcome(
		results,
		s.store.CollectionFacets(collection, metadata, 20),
	), nil
}
