package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/symbol_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestTheClientRoundTripsARetarget(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	reports, e := c.PostReferencesRetarget(
		t.Context(),
		move_set.New(
			[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
			[]*symbol_move.Move{
				symbol_move.New("pkg/alfa", "Old", "pkg/alfa", "New"),
			},
		),
		true,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, reports)
}
