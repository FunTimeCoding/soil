package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/filter"
)

func (s *Service) Expiring() ([]record.Record, error) {
	horizon := filter.New()
	horizon.Window.WithEnd(s.clock().AddDate(0, 0, constant.ExpiryHorizonDay))
	horizon.Revoked = new(false)

	return s.store.Certificates(horizon)
}
