package service

import "github.com/prometheus/alertmanager/api/v2/models"

func (s *Service) Status(instance string) (*models.AlertmanagerStatus, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Status()
}
