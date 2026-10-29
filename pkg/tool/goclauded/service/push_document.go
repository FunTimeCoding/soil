package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) pushDocument(
	collection string,
	sourceType string,
	path string,
	body string,
	metadata map[string][]string,
) error {
	_, e := s.query.PostDocument(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: collection,
			Path:       path,
			Body:       body,
			SourceType: &sourceType,
			Metadata:   &metadata,
		},
	)

	return e
}
