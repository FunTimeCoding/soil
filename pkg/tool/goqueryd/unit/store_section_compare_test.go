package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/comparison"
	"testing"
)

func TestCompareRestructureKeepsEveryWord(t *testing.T) {
	c := comparison.Compare(
		constant.FixtureCompareBefore,
		constant.FixtureCompareRestructured,
	)
	assert.Integer(t, 0, c.Removed)
	assert.Integer(t, 0, c.Added)
	assert.Count(t, 2, c.Changes)
	assert.String(t, "# A", c.Changes[0].Title)
	assert.Integer(t, 2, c.Changes[0].RemovedCount)
	assert.String(t, "## B", c.Changes[1].Title)
	assert.Integer(t, 2, c.Changes[1].AddedCount)
}

func TestCompareCatchesOneShavedWord(t *testing.T) {
	c := comparison.Compare(
		constant.FixtureCompareBefore,
		constant.FixtureCompareShaved,
	)
	assert.Integer(t, 1, c.Removed)
	assert.Integer(t, 0, c.Added)
	assert.Strings(t, []string{"bravo"}, c.Changes[0].Removed)
}
