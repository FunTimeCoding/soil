package service

import "context"

func (s *Service) Wait(x context.Context) error {
	if e := s.jiraGuard.Wait(x); e != nil {
		return e
	}

	return s.confluenceGuard.Wait(x)
}
