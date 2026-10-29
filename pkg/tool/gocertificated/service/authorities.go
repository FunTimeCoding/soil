package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Service) Authorities() ([]record.Record, error) {
	return s.store.Authorities()
}
