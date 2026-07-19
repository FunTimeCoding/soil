package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/lower"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
	"time"
)

func TestRecentSeedsOrderNewestFirst(t *testing.T) {
	s := store_tester.New(t)
	base := time.Now()
	s.Store.UpsertSeed(
		lower.Alfa,
		"alfa.md",
		"hash-a",
		"a",
		base.Add(-2*time.Hour),
	)
	s.Store.UpsertSeed(lower.Bravo, "bravo.md", "hash-b", "b", base)
	s.Store.UpsertSeed(
		lower.Charlie,
		"charlie.md",
		"hash-c",
		"c",
		base.Add(-time.Hour),
	)
	seeds := s.Store.RecentSeeds()
	assert.Count(t, 3, seeds)
	assert.String(t, "bravo", seeds[0].Name)
	assert.String(t, "charlie", seeds[1].Name)
	assert.String(t, "alfa", seeds[2].Name)
}
