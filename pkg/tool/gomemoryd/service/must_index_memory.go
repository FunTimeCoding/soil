package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) MustIndexMemory(
	path string,
	content string,
	metadata map[string]string,
) {
	errors.PanicOnError(s.IndexMemory(path, content, metadata))
}
