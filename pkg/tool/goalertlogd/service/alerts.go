package service

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"

func (s *Service) Alerts(name string) ([]record.Record, error) {
	return s.store.ByName(name)
}
