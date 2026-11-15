package store_tester

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"time"
)

func ThreeSeeds(s *Tester) {
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
