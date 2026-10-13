package service

import "time"

func (s *Service) CreateSilence(
	instance string,
	alert string,
	comment string,
	d time.Duration,
) (string, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return "", e
	}

	return c.SetSilence(alert, comment, d)
}
