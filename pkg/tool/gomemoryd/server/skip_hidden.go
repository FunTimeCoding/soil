package server

import "slices"

func (s *Server) skipHidden(tags []string) bool {
	tag := s.hiddenTag

	return tag != "" && slices.Contains(tags, tag)
}
