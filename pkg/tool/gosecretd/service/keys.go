package service

import (
	"maps"
	"slices"
)

func (s *Service) Keys(path string) ([]string, error) {
	if e := validatePath(path); e != nil {
		return nil, e
	}

	values, e := s.source.Read(path)

	if e != nil {
		return nil, e
	}

	return slices.Sorted(maps.Keys(values)), nil
}
