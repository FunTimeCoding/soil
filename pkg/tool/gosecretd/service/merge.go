package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"maps"
	"slices"
)

func (s *Service) Merge(
	path string,
	values map[string]string,
) ([]string, error) {
	if e := validatePath(path); e != nil {
		return nil, e
	}

	if len(values) == 0 {
		return nil, validation.New("values are required")
	}

	result, e := s.source.Read(path)

	if not_found.Is(e) {
		result = map[string]string{}
	} else if e != nil {
		return nil, e
	}

	maps.Copy(result, values)

	if f := s.source.Put(path, result); f != nil {
		return nil, f
	}

	return slices.Sorted(maps.Keys(result)), nil
}
