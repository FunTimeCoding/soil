package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"maps"
	"slices"
)

func (s *Service) Keys(path string) ([]string, error) {
	if e := validatePath(path); e != nil {
		return nil, e
	}

	values, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (map[string]string, error) {
			return s.client.Read(path)
		},
	)

	if e != nil {
		return nil, e
	}

	return slices.Sorted(maps.Keys(values)), nil
}
