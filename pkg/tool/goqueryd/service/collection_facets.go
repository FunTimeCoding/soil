package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func (s *Service) CollectionFacets(
	collection string,
	metadata map[string]string,
) []search.Facet {
	return s.store.CollectionFacets(collection, metadata, 20)
}
