//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func TestUpsertNewSeed(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "content a", time.Now())
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "alfa.md", seeds[0].Path)
	assert.Integer(t, 1, seeds[0].Position)
}

func TestUpsertSecondSeedGetsNextPosition(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "content a", time.Now())
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "content b", time.Now())
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
	assert.Integer(t, 1, seeds[0].Position)
	assert.Integer(t, 2, seeds[1].Position)
}

func TestUpsertUpdatesContentOnHashChange(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "old content", time.Now())
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-2", "new content", time.Now())
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
	assert.String(t, "hash-2", seeds[0].ContentHash)
	assert.String(t, "new content", seeds[0].Content)
}

func TestUpsertNoOpOnSameHash(t *testing.T) {
	s := store_tester.New(t)
	modified := time.Now()
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "content", modified)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "content", modified)
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
}

func TestUpsertUpdatesModifiedTimeWithoutContentChange(t *testing.T) {
	s := store_tester.New(t)
	base := time.Now()
	s.Store.UpsertSeed(
		lower.Alfa,
		"alfa.md",
		"hash-1",
		"content",
		base.Add(-time.Hour),
	)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-1", "content", base)
	seeds := s.Store.Seeds()
	assert.Count(t, 1, seeds)
	assert.True(t, seeds[0].ModifiedAt.After(base.Add(-time.Minute)))
}

func TestUpsertSameNameDifferentPath(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "root", time.Now())
	s.Store.UpsertSeed(lower.Alfa, "sub/alfa.md", "hash-b", "nested", time.Now())
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
}
