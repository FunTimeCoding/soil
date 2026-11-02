package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func TestRecentSeedsOrderNewestFirst(t *testing.T) {
	s := store_tester.New(t)
	base := time.Now()
	s.Store.UpsertSeed(
		stringsConstant.LowerAlfa,
		"alfa.md",
		"hash-a",
		"a",
		base.Add(-2*time.Hour),
	)
	s.Store.UpsertSeed(
		stringsConstant.LowerBravo,
		"bravo.md",
		"hash-b",
		"b",
		base,
	)
	s.Store.UpsertSeed(
		stringsConstant.LowerCharlie,
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
