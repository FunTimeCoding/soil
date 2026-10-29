package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"time"
)

func (s *Service) Revoke(serial string) (*record.Record, error) {
	existing, e := s.store.BySerial(serial)

	if e != nil {
		return nil, e
	}

	if existing.Revoked != nil {
		return nil, not_found.New(constant.SerialParameter, serial)
	}

	if f := s.store.Revoke(serial, time.Now()); f != nil {
		return nil, f
	}

	return s.store.BySerial(serial)
}
