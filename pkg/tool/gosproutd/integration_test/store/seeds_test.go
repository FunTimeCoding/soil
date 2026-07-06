//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestSeedsOrderedByPosition(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Charlie, "charlie.md", "hash-c", "c")
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "a")
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "b")
	seeds := s.Store.Seeds()
	assert.String(t, "charlie", seeds[0].Name)
	assert.String(t, "alfa", seeds[1].Name)
	assert.String(t, "bravo", seeds[2].Name)
}

func TestSeedsEmptyStore(t *testing.T) {
	s := store_tester.New(t)
	seeds := s.Store.Seeds()
	assert.Count(t, 0, seeds)
}
