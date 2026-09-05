package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/integration/store_tester"
	"testing"
	"time"
)

func TestRecentSeedsOrderNewestFirst(t *testing.T) {
	s := store_tester.New(t)
	base := time.Now()
	s.Store.UpsertSeed(
		constant.LowerAlfa,
		"alfa.md",
		"hash-a",
		"a",
		base.Add(-2*time.Hour),
	)
	s.Store.UpsertSeed(constant.LowerBravo, "bravo.md", "hash-b", "b", base)
	s.Store.UpsertSeed(
		constant.LowerCharlie,
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
