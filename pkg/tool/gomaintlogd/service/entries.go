package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/filter"
)

func (s *Service) Entries(f *filter.Filter) ([]entry.Entry, error) {
	return s.store.List(f)
}
