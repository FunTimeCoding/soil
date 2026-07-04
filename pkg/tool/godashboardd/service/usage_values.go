package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"strconv"
)

func (s *Service) usageValues() []string {
	if s.usage == nil {
		return []string{constant.PendingValue, constant.PendingValue}
	}

	result, e := s.usage.Fetch()

	if e != nil {
		s.logger.Plain("usage fetch failed: %v", e)

		return []string{constant.PendingValue, constant.PendingValue}
	}

	return []string{
		strconv.FormatInt(result.Files, 10),
		strconv.FormatInt(result.Shares, 10),
	}
}
