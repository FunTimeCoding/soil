package service

import "context"

func (s *Service) Warm(x context.Context) []error {
	for _, e := range s.Sessions() {
		if x.Err() != nil {
			return nil
		}

		s.ToolCalls(e.Identifier)
	}

	return nil
}
