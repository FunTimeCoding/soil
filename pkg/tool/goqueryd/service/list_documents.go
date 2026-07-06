package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func (s *Service) ListDocuments(
	collection string,
	metadata map[string]string,
	limit int,
	offset int,
	full bool,
) (*store.ListOutcome, error) {
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

	return store.NewListOutcome(
		results,
		s.store.CollectionFacets(collection, metadata, 20),
	), nil
}
