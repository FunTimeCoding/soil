package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosource"
	"testing"
)

func TestPruneLinesNameEachRemovedGeneration(t *testing.T) {
	assert.Strings(
		t,
		[]string{"removed go1.0-stale", "removed go1.1-old"},
		gosource.PruneLines([]string{"go1.0-stale", "go1.1-old"}),
	)
}

func TestPruneLinesSayWhenNothingWasStale(t *testing.T) {
	assert.Strings(t, []string{"nothing to prune"}, gosource.PruneLines(nil))
}
