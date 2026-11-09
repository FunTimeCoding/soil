package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestClientRoundTripsRetarget(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "pkg/bravo/run.go", "package bravo\n")
	c := base.NewRooted(t, root).MemoryClient(t)
	_, e := c.PostMemory(
		t.Context(),
		client.PostMemoryJSONRequestBody{
			Name:        "plain",
			Content:     "Read `pkg/alfa/run.go` first.",
			Description: "plain",
		},
	)
	assert.FatalOnError(t, e)
	reports, f := c.PostReferencesRetarget(
		t.Context(),
		move_set.New(
			[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
			nil,
		),
		true,
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, reports)
	assert.String(t, "plain", reports[0].Name)
	assert.String(t, "pkg/bravo/run.go", reports[0].Changes[0].New)
}
