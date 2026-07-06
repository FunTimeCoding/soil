//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestSetPositionMoveToTop(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.SetPosition(seeds[2].Identifier, 1)
	after := s.Store.Seeds()
	assert.String(t, "charlie", after[0].Name)
	assert.String(t, "alfa", after[1].Name)
	assert.String(t, "bravo", after[2].Name)
}

func TestSetPositionMoveToBottom(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.SetPosition(seeds[0].Identifier, 3)
	after := s.Store.Seeds()
	assert.String(t, "bravo", after[0].Name)
	assert.String(t, "charlie", after[1].Name)
	assert.String(t, "alfa", after[2].Name)
}

func TestSetPositionMoveToMiddle(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.SetPosition(seeds[2].Identifier, 2)
	after := s.Store.Seeds()
	assert.String(t, "alfa", after[0].Name)
	assert.String(t, "charlie", after[1].Name)
	assert.String(t, "bravo", after[2].Name)
}

func TestSetPositionSamePositionIsNoOp(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.SetPosition(seeds[1].Identifier, 2)
	after := s.Store.Seeds()
	assert.String(t, "alfa", after[0].Name)
	assert.String(t, "bravo", after[1].Name)
	assert.String(t, "charlie", after[2].Name)
}
