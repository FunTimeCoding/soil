package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"math"
)

func (s *Service) CollectionFacetsForKey(
	collection string,
	key string,
) []result.Facet {
	return s.store.CollectionFacets(collection, nil, math.MaxInt, key)
}
