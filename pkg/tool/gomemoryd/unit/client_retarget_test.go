package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestARetargetWithoutAReferenceRootSaysSo(t *testing.T) {
	c := base.New(t).MemoryClient(t)
	_, e := c.PostReferencesRetarget(
		t.Context(),
		move_set.New(
			[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
			nil,
		),
		true,
	)
	assert.StringContains(
		t,
		"GOMEMORY_REFERENCE_ROOT is not set - no reference root, nothing checked",
		e.Error(),
	)
}
