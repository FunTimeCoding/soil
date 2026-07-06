package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (s *Service) QueryRange(
	instance string,
	query string,
	r v1.Range,
) (*query_result.Result, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.QueryRange(query, r)
}
