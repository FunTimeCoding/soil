//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	ollamaConstant "github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	goqueryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"os"
	"path/filepath"
	"testing"
)

func openTestStore(t *testing.T) (*Store, *ollama.Client) {
	t.Helper()

	return New(connection.NewMemory()), ollama.NewEnvironment()
}

func writeFixture(
	t *testing.T,
	directory string,
	name string,
	content string,
) {
	t.Helper()
	path := filepath.Join(directory, name)
	errors.PanicOnError(os.MkdirAll(filepath.Dir(path), 0o755))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o644))
}

func indexedTestStore(t *testing.T) (*Store, *ollama.Client) {
	t.Helper()
	s, o := openTestStore(t)
	s.AddCollection(
		"test",
		fixture.Path(system.SearchPath),
		goqueryd.DefaultGlob,
	)
	s.Index("test")

	return s, o
}

func pushTestDocument(
	s *Store,
	o *ollama.Client,
	collection string,
	path string,
	body string,
	sourceType string,
) error {
	s.EnsurePushCollection(collection)
	now := "2024-01-01T00:00:00Z"
	hash := HashContent(body)
	title := ExtractTitle(body, path)
	s.InsertContent(hash, body, now)
	s.InsertDocument(collection, path, title, hash, now)
	chunks := chunk.Document(body, path)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := o.Embed(ollamaConstant.EmbedModel, texts)

	if e != nil {
		return e
	}

	for i, embedding := range embeddings {
		s.InsertEmbedding(hash, i, chunks[i].Position, embedding, now)
	}

	if sourceType != "" {
		s.SetSourceType(collection, path, sourceType)
	}

	return nil
}

func embedTestDocuments(
	s *Store,
	o *ollama.Client,
) error {
	pending := s.PendingEmbeddings()
	now := "2024-01-01T00:00:00Z"

	for _, p := range pending {
		chunks := chunk.Document(p.Body, p.Path)
		texts := make([]string, len(chunks))

		for i, c := range chunks {
			texts[i] = c.Text
		}

		embeddings, e := o.Embed(ollamaConstant.EmbedModel, texts)

		if e != nil {
			return e
		}

		for i, embedding := range embeddings {
			s.InsertEmbedding(
				p.Hash,
				i,
				chunks[i].Position,
				embedding,
				now,
			)
		}
	}

	return nil
}
