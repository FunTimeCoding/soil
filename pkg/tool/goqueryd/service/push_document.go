package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

func (s *Service) PushDocument(
	x context.Context,
	collection string,
	path string,
	body string,
	metadata map[string][]string,
) string {
	s.store.EnsurePushCollection(collection)
	now := time.Now().UTC().Format(time.RFC3339)
	hash := store.HashContent(body)
	fresh := s.store.ActiveDocumentHash(collection, path) != hash

	if fresh {
		s.store.InsertContent(hash, body, now)
		s.store.InsertDocument(
			collection,
			path,
			store.ExtractTitle(body, path),
			hash,
			now,
		)
	}

	if len(metadata) > 0 {
		s.store.SetMetadata(collection, path, metadata)
	}

	if !fresh {
		return ""
	}

	_, e := s.embedDocument(x, hash, path, body)

	switch {
	case e == nil:
		return ""
	case validation.Is(e):
		return fmt.Sprintf(constant.EmbeddingRefused, e)
	default:
		return fmt.Sprintf(constant.EmbeddingPending, e)
	}
}
