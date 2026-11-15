package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/line_diff"
	"github.com/funtimecoding/soil/pkg/text/unit/line_diff_tester"
	"testing"
)

func TestLineDiffAnchorsInsertedLineToLineItFollows(t *testing.T) {
	anchors := line_diff.Anchors(
		"alfa\nbravo\ncharlie\n",
		"alfa\nbravo\n\ncharlie\n",
	)
	assert.Integers(t, []int{1, 2, 2, 3}, line_diff_tester.Lines(anchors))
	assert.False(t, anchors[1].Rewritten)
	assert.True(t, anchors[2].Rewritten)
	assert.False(t, anchors[3].Rewritten)
}

func TestLineDiffAnchorsReplacementToFirstLineReplaced(t *testing.T) {
	anchors := line_diff.Anchors(
		"alfa\nbravo charlie\ndelta\n",
		"alfa\nbravo\ncharlie\ndelta\n",
	)
	assert.Integers(t, []int{1, 2, 2, 3}, line_diff_tester.Lines(anchors))
	assert.True(t, anchors[1].Rewritten)
	assert.True(t, anchors[2].Rewritten)
	assert.False(t, anchors[3].Rewritten)
}
