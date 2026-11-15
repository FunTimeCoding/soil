package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func Push(
	s *store.Store,
	o *ollama.Client,
	collection string,
	path string,
	body string,
	sourceType string,
) error {
	if e := s.EnsurePushCollection(collection); e != nil {
		return e
	}

	now := "2024-01-01T00:00:00Z"
	hash := store.HashContent(body)
	title := store.ExtractTitle(body, path)

	if e := s.InsertContent(hash, body, now); e != nil {
		return e
	}

	if e := s.InsertDocument(collection, path, title, hash, now); e != nil {
		return e
	}

	if e := insertEmbeddings(s, o, hash, path, body, now); e != nil {
		return e
	}

	if sourceType != "" {
		s.SetSourceType(collection, path, sourceType)
	}

	return nil
}
