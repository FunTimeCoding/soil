package score

import "context"

func (s *Service) Sync(x context.Context) []error {
	if e := s.load(); e != nil {
		return []error{e}
	}

	s.mutex.RLock()
	m := *s.state
	s.mutex.RUnlock()
	now := s.clock()

	if !m.Filled {
		if e := s.fill(x, &m, now); e != nil {
			return []error{e}
		}

		return nil
	}

	if e := s.update(x, &m, now); e != nil {
		return []error{e}
	}

	return nil
}
