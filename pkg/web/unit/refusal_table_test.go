package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
	"slices"
	"testing"
)

func TestWithoutDropsTheRowAndLeavesTheTableItCameFrom(t *testing.T) {
	table := refusal.NewTable()
	assert.Nil(
		t,
		table.Without(http.StatusServiceUnavailable).Classify(
			http.StatusServiceUnavailable,
			"alfa",
		),
	)
	assert.String(t, "unavailable", table.Class(http.StatusServiceUnavailable))
}

func TestEveryRowNamesAClassThatClassedKnows(t *testing.T) {
	known := []string{
		errorsConstant.ValidationClass,
		errorsConstant.NotFoundClass,
		errorsConstant.ConflictClass,
		errorsConstant.ForbiddenClass,
		errorsConstant.UnavailableClass,
	}

	for _, table := range []*refusal.Table{
		refusal.NewTable(),
		refusal.NewDaemonTable(),
	} {
		for _, class := range table.Rows() {
			assert.True(t, slices.Contains(known, class))
		}
	}
}
