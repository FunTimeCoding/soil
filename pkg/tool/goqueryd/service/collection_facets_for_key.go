package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"math"
)

func (s *Service) CollectionFacetsForKey(
	collection string,
	key string,
) []store.Facet {
	return s.store.CollectionFacets(
		collection,
		nil,
		math.MaxInt,
		key,
	)
}
