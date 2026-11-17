package store_tester

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"time"
)

func (o *Tester) ThreeSeeds() {
	o.Store.UpsertSeed(
		stringsConstant.LowerAlfa,
		"alfa.md",
		"hash-a",
		"a",
		time.Now(),
	)
	o.Store.UpsertSeed(
		stringsConstant.LowerBravo,
		"bravo.md",
		"hash-b",
		"b",
		time.Now(),
	)
	o.Store.UpsertSeed(
		stringsConstant.LowerCharlie,
		"charlie.md",
		"hash-c",
		"c",
		time.Now(),
	)
}
