package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) unindex(
	scope string,
	identifier int64,
) error {
	_, e := s.query.DeleteDocument(
		context.Background(),
		&client.DeleteDocumentParams{
			Collection: ScopeCollection(scope),
			Path:       memoryPath(identifier),
		},
	)

	return e
}
