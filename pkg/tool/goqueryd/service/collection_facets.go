package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func (s *Service) CollectionFacets(
	collection string,
	metadata map[string]string,
) []result.Facet {
	return s.store.CollectionFacets(collection, metadata, 20)
}
