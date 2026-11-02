package unit

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"time"
)

func threeSeeds(s *store_tester.Tester) {
	s.Store.UpsertSeed(
		stringsConstant.LowerAlfa,
		"alfa.md",
		"hash-a",
		"a",
		time.Now(),
	)
	s.Store.UpsertSeed(
		stringsConstant.LowerBravo,
		"bravo.md",
		"hash-b",
		"b",
		time.Now(),
	)
	s.Store.UpsertSeed(
		stringsConstant.LowerCharlie,
		"charlie.md",
		"hash-c",
		"c",
		time.Now(),
	)
}
