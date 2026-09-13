package service

import "github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"

func (s *Service) nextGroupNumber() (int, error) {
	found, e := s.Groups()

	if e != nil {
		return 0, e
	}

	result := constant.FirstGroupNumber

	for _, g := range found {
		if g.Number >= result {
			result = g.Number + 1
		}
	}

	return result, nil
}
