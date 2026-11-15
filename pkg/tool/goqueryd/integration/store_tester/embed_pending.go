package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func EmbedPending(
	s *store.Store,
	o *ollama.Client,
) error {
	now := "2024-01-01T00:00:00Z"

	for _, p := range s.PendingEmbeddings() {
		if e := insertEmbeddings(s, o, p.Hash, p.Path, p.Body, now); e != nil {
			return e
		}
	}

	return nil
}
