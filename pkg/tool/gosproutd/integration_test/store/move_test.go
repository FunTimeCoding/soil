//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func threeSeeds(s *store_tester.Tester) {
	s.Store.UpsertSeed(lower.Alfa, "alfa.md", "hash-a", "a", time.Now())
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "b", time.Now())
	s.Store.UpsertSeed(lower.Charlie, "charlie.md", "hash-c", "c", time.Now())
}

func TestMoveUpSwapsWithAbove(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveUp(seeds[1].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "bravo", after[0].Name)
	assert.String(t, "alfa", after[1].Name)
	assert.String(t, "charlie", after[2].Name)
}

func TestMoveUpAtTopIsNoOp(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveUp(seeds[0].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "alfa", after[0].Name)
}

func TestMoveDownSwapsWithBelow(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveDown(seeds[0].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "bravo", after[0].Name)
	assert.String(t, "alfa", after[1].Name)
	assert.String(t, "charlie", after[2].Name)
}

func TestMoveDownAtBottomIsNoOp(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveDown(seeds[2].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "charlie", after[2].Name)
}
