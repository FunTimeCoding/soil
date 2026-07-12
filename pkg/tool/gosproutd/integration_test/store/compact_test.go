//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func TestCompactClosesGaps(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "a", time.Now())
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "b", time.Now())
	s.Store.UpsertSeed(lower.Charlie, "charlie.md", "hash-c", "c", time.Now())
	s.Store.RemoveMissing([]string{"alfa.md", "charlie.md"})
	s.Store.Compact()
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
	assert.Integer(t, 1, seeds[0].Position)
	assert.Integer(t, 2, seeds[1].Position)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "charlie", seeds[1].Name)
}

func TestCompactNoOpWhenSequential(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	s.Store.Compact()
	seeds := s.Store.Seeds()
	assert.Integer(t, 1, seeds[0].Position)
	assert.Integer(t, 2, seeds[1].Position)
	assert.Integer(t, 3, seeds[2].Position)
}

func TestCompactEmptyStore(t *testing.T) {
	s := store_tester.New(t)
	s.Store.Compact()
	seeds := s.Store.Seeds()
	assert.Count(t, 0, seeds)
}

func TestCompactPreservesOrder(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "a", time.Now())
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "b", time.Now())
	s.Store.UpsertSeed(lower.Charlie, "charlie.md", "hash-c", "c", time.Now())
	s.Store.UpsertSeed(lower.Delta, "delta.md", "hash-d", "d", time.Now())
	s.Store.RemoveMissing([]string{"alfa.md", "delta.md"})
	s.Store.Compact()
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "delta", seeds[1].Name)
	assert.Integer(t, 1, seeds[0].Position)
	assert.Integer(t, 2, seeds[1].Position)
}
