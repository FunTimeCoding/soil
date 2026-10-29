package service

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Service) BySerial(serial string) (*record.Record, error) {
	return s.store.BySerial(serial)
}
