package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"strconv"
)

func (s *Service) usageValues() ([]string, error) {
	pending := []string{constant.PendingValue, constant.PendingValue}

	if s.usage == nil {
		return pending, nil
	}

	result, e := guarded(s.usageGuard, s.usage.Fetch)

	if e != nil {
		return pending, e
	}

	return []string{
		strconv.FormatInt(result.Files, 10),
		strconv.FormatInt(result.Shares, 10),
	}, nil
}
