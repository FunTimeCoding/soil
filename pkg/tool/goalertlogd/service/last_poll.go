package service

import "time"

func (s *Service) LastPoll() time.Time {
	v := s.lastPoll.Load()

	if v == nil {
		return time.Time{}
	}

	return v.(time.Time)
}
