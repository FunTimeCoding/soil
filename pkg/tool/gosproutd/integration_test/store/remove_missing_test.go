//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestRemoveMissingDeletesAbsentPaths(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	s.Store.RemoveMissing([]string{"alfa.md", "charlie.md"})
	seeds := s.Store.Seeds()
	assert.Count(t, 2, seeds)
	assert.String(t, "alfa", seeds[0].Name)
	assert.String(t, "charlie", seeds[1].Name)
}

func TestRemoveMissingEmptyListClearsAll(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	s.Store.RemoveMissing([]string{})
	seeds := s.Store.Seeds()
	assert.Count(t, 0, seeds)
}

func TestRemoveMissingNilClearsAll(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	s.Store.RemoveMissing(nil)
	seeds := s.Store.Seeds()
	assert.Count(t, 0, seeds)
}
