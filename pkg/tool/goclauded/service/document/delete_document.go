package document

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) deleteDocument(
	collection string,
	path string,
) error {
	return s.queryGuard.Attempt(
		context.Background(),
		func(x context.Context) error {
			_, e := s.query.DeleteDocument(
				x,
				&client.DeleteDocumentParams{Collection: collection, Path: path},
			)

			return e
		},
	)
}
