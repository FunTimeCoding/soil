package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/comparison"
)

func compareWords(
	before string,
	after string,
) int {
	c := comparison.Compare(before, after)
	console.Format("\ncompared with snapshot\n")

	for _, x := range c.Changes {
		console.Format(
			"  %s: removed %d, added %d\n",
			label(x.Title),
			x.RemovedCount,
			x.AddedCount,
		)
		printWords("removed", x.Removed)
		printWords("added", x.Added)
	}

	console.Format(
		"  overall: removed %d, added %d words\n",
		c.Removed,
		c.Added,
	)

	return c.Removed + c.Added
}
