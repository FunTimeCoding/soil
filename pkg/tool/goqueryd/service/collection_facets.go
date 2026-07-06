package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func (s *Service) CollectionFacets(
	collection string,
	metadata map[string]string,
) []store.Facet {
	return s.store.CollectionFacets(collection, metadata, 20)
}
