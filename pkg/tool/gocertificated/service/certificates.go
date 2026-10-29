package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/filter"
)

func (s *Service) Certificates(f *filter.Filter) ([]record.Record, error) {
	return s.store.Certificates(f)
}
