package service

import "slices"

func (s *Service) isHidden(tags []string) bool {
	return s.hiddenTag != "" &&
		slices.Contains(tags, s.hiddenTag)
}
