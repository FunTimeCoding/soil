package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func TestSeedsOrderedByPosition(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed(
		constant.LowerCharlie,
		"charlie.md",
		"hash-c",
		"c",
		time.Now(),
	)
	s.Store.UpsertSeed(
		constant.LowerAlfa,
		"alfa.md",
		"hash-a",
		"a",
		time.Now(),
	)
	s.Store.UpsertSeed(
		constant.LowerBravo,
		"bravo.md",
		"hash-b",
		"b",
		time.Now(),
	)
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
