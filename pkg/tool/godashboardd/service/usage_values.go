package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"strconv"
)

func (s *Service) usageValues() ([]string, error) {
	pending := []string{constant.PendingValue, constant.PendingValue}

	if s.usage == nil {
		return pending, nil
	}

	result, e := peer.Guarded(
		context.Background(),
		s.usageGuard,
		func(context.Context) (*usage.Usage, error) { return s.usage.Fetch() },
	)

	if e != nil {
		return pending, e
	}

	return []string{
		strconv.FormatInt(result.Files, 10),
		strconv.FormatInt(result.Shares, 10),
	}, nil
}
