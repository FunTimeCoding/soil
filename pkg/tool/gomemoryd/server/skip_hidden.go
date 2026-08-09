package server

import "slices"

func (s *Server) skipHidden(tags []string) bool {
	tag := s.service.HiddenTag()

	return tag != "" && slices.Contains(tags, tag)
}
