package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/comparison"
)

func compareWords(
	t *terminal.Terminal,
	before string,
	after string,
) int {
	c := comparison.Compare(before, after)
	t.Line()
	t.Line("compared with snapshot")
	cut := 0

	for _, x := range c.Changes {
		t.Line(
			fmt.Sprintf(
				"  %s: removed %d, added %d",
				label(x.Title),
				x.RemovedCount,
				x.AddedCount,
			),
		)
		removed := printWords(t, "removed", x.Removed)

		if printWords(t, "added", x.Added) || removed {
			cut++
		}
	}

	t.Line(
		fmt.Sprintf(
			"  overall: removed %d, added %d words",
			c.Removed,
			c.Added,
		),
	)

	if cut > 0 {
		t.Noticeln(
			fmt.Sprintf(
				"word lists cut at %d in %d sections",
				constant.WordListLimit,
				cut,
			),
		)
	}

	return c.Removed + c.Added
}
