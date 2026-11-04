package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (s *Service) MustPushDocument(
	x context.Context,
	collection string,
	path string,
	body string,
	metadata map[string][]string,
) string {
	result, e := s.PushDocument(x, collection, path, body, metadata)
	errors.PanicOnError(e)

	return result
}
