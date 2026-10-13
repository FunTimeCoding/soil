package service

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (s *Service) Series(instance string) (v1.TSDBResult, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return v1.TSDBResult{}, e
	}

	return c.Series()
}
