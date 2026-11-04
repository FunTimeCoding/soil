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
) (string, error) {
	if e := s.store.EnsurePushCollection(collection); e != nil {
		return "", classifyWrite(e)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	hash := store.HashContent(body)
	fresh := s.store.ActiveDocumentHash(collection, path) != hash

	if fresh {
		if e := s.store.InsertContent(hash, body, now); e != nil {
			return "", classifyWrite(e)
		}

		if e := s.store.InsertDocument(
			collection,
			path,
			store.ExtractTitle(body, path),
			hash,
			now,
		); e != nil {
			return "", classifyWrite(e)
		}
	}

	if len(metadata) > 0 {
		if e := s.store.SetMetadata(collection, path, metadata); e != nil {
			return "", classifyWrite(e)
		}
	}

	if !fresh {
		return "", nil
	}

	_, e := s.embedDocument(x, hash, path, body)

	switch {
	case e == nil:
		return "", nil
	case validation.Is(e):
		return fmt.Sprintf(constant.EmbeddingRefused, e), nil
	default:
		return fmt.Sprintf(constant.EmbeddingPending, e), nil
	}
}
