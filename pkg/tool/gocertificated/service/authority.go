package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Service) Authority(name string) (*record.Record, error) {
	return s.store.Authority(name)
}
