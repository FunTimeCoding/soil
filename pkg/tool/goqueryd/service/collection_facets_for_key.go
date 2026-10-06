package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"math"
)

func (s *Service) CollectionFacetsForKey(
	collection string,
	key string,
) []search.Facet {
	return s.store.CollectionFacets(collection, nil, math.MaxInt, key)
}
