package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"time"
)

func (s *Service) queryValue(q string) string {
	result, e := s.prometheus.Query(q, time.Now())

	if e != nil {
		s.logger.Plain("query failed: %s: %v", q, e)

		return constant.PendingValue
	}

	return formatValue(result.Value)
}
