package search

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
)

func New(
	s *store.Store,
	x *search_index.Index,
	harbor string,
) *Service {
	return &Service{store: s, search: x, harbor: harbor}
}
