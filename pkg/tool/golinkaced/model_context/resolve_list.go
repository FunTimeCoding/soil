package model_context

import "strconv"

func (s *Server) resolveList(value string) (int, error) {
	if value == "" {
		return 0, nil
	}

	identifier, e := strconv.Atoi(value)

	if e == nil {
		return identifier, nil
	}

	l, f := s.client.ListByName(value)

	if f != nil {
		return 0, f
	}

	return l.Identifier, nil
}
