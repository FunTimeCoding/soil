package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

func (s *Service) Query(
	instance string,
	query string,
	t time.Time,
) (*query_result.Result, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Query(query, t)
}
