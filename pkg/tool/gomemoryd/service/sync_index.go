package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) syncIndex(m *record.Memory) error {
	if !s.indexable(m) {
		return s.unindex(m.Scope, m.Identifier)
	}

	metadata := pushMetadata(m)
	sourceType := constant.MemorySourceType
	_, e := s.query.PostDocument(
		context.Background(),
		client.PostDocumentJSONRequestBody{
			Collection: ScopeCollection(m.Scope),
			Path:       memoryPath(m.Identifier),
			Body:       m.Content,
			SourceType: &sourceType,
			Metadata:   &metadata,
		},
	)

	return e
}
