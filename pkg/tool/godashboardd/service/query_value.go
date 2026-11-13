package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"time"
)

func (s *Service) queryValue(q string) (string, error) {
	r, e := peer.Guarded(
		context.Background(),
		s.prometheusGuard,
		func(context.Context) (*query_result.Result, error) {
			return s.prometheus.Query(q, time.Now())
		},
	)

	if e != nil {
		return constant.PendingValue, fmt.Errorf("query %s: %w", q, e)
	}

	return formatValue(r.Value), nil
}
