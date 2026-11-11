package cruise

import "time"

func (s *Service) WithDefaultGrace(d time.Duration) *Service {
	s.defaultGrace = d

	return s
}
