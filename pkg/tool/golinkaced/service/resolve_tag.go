package service

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (s *Service) resolveTag(name string) (int, error) {
	t, e := guarded(
		s.guard,
		func() (*tag.Tag, error) { return s.client.TagByName(name) },
	)

	if e != nil {
		return 0, e
	}

	return t.Identifier, nil
}
