package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) syncIndex(
	x context.Context,
	m *record.Memory,
) error {
	if !s.indexable(m) {
		return s.unindex(x, m.Scope, m.Identifier)
	}

	metadata := pushMetadata(m)
	sourceType := constant.MemorySourceType

	return s.guard.Attempt(
		x,
		func(y context.Context) error {
			_, e := s.query.PostDocument(
				y,
				client.PostDocumentJSONRequestBody{
					Collection: ScopeCollection(m.Scope),
					Path:       memoryPath(m.Identifier),
					Body:       m.Content,
					SourceType: &sourceType,
					Metadata:   &metadata,
				},
			)

			return e
		},
	)
}
