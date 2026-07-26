package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func threeSeeds(s *store_tester.Tester) {
	s.Store.UpsertSeed(constant.LowerAlfa, "alfa.md", "hash-a", "a", time.Now())
	s.Store.UpsertSeed(constant.LowerBravo, "bravo.md", "hash-b", "b", time.Now())
	s.Store.UpsertSeed(
		constant.LowerCharlie,
		"charlie.md",
		"hash-c",
		"c",
		time.Now(),
	)
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
