package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/line_diff"
	"testing"
)

func TestLineDiffShowsChangesWithContext(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"@@",
			" charlie",
			" delta",
			"-echo",
			"+foxtrot",
			" golf",
			" hotel",
		},
		line_diff.Hunks(
			line_diff.Lines(
				"alfa\nbravo\ncharlie\ndelta\necho\ngolf\nhotel\nindia\njuliett\n",
				"alfa\nbravo\ncharlie\ndelta\nfoxtrot\ngolf\nhotel\nindia\njuliett\n",
			),
			2,
		),
	)
}

func TestLineDiffOfEqualTextIsEmpty(t *testing.T) {
	assert.Count(t, 0, line_diff.Hunks(line_diff.Lines("alfa\n", "alfa\n"), 2))
}
