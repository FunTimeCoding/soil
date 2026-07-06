package scan

import "github.com/funtimecoding/soil/pkg/lint/concern"

func (s *Service) addConcern(
	key string,
	text string,
	path string,
) {
	s.Concerns = append(
		s.Concerns,
		concern.NewPackage(key, text, path),
	)
}
