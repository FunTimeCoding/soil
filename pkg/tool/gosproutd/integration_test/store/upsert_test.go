//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestUpsertNewSeed(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "content a")
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "alfa.md", seeds[0].Path)
	assert.Integer(t, 1, seeds[0].Position)
}

func TestUpsertSecondSeedGetsNextPosition(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "content a")
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "content b")
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
	assert.Integer(t, 1, seeds[0].Position)
	assert.Integer(t, 2, seeds[1].Position)
}

func TestUpsertUpdatesContentOnHashChange(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "old content")
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-2", "new content")
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
	assert.String(t, "hash-2", seeds[0].ContentHash)
	assert.String(t, "new content", seeds[0].Content)
}

func TestUpsertNoOpOnSameHash(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "content")
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "content")
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
}

func TestUpsertSameNameDifferentPath(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "root")
	s.Store.UpsertSeed(lower.Alfa, "sub/alfa.md", "hash-b", "nested")
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
}
