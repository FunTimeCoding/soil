package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func (s *Service) QueryRange(
	x context.Context,
	instance string,
	query string,
	start string,
	end string,
	step time.Duration,
) (*query_result.Result, error) {
	now := s.clock()
	from, e := parseMoment(start, now)

	if e != nil {
		return nil, validation.New("invalid start: %s", e)
	}

	to := now

	if end != "" {
		if to, e = parseMoment(end, now); e != nil {
			return nil, validation.New("invalid end: %s", e)
		}
	}

	return guarded(
		s,
		x,
		instance,
		func(c face.PrometheusClient) (*query_result.Result, error) {
			return c.QueryRange(
				query,
				v1.Range{Start: from, End: to, Step: step},
			)
		},
	)
}
