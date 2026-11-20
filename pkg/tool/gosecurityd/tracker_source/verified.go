package tracker_source

import "time"

func (s *Source) Verified() time.Time {
	if v := s.verified.Load(); v != nil {
		return *v
	}

	return time.Time{}
}
