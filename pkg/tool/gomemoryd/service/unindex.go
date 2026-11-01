package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) unindex(
	x context.Context,
	scope string,
	identifier int64,
) error {
	return s.guard.Attempt(
		x,
		func(y context.Context) error {
			_, e := s.query.DeleteDocument(
				y,
				&client.DeleteDocumentParams{
					Collection: ScopeCollection(scope),
					Path:       memoryPath(identifier),
				},
			)

			return e
		},
	)
}
