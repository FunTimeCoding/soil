package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestRetargetTreeRewritesOnlyWhatMentionsTheMove(t *testing.T) {
	d := t.TempDir()
	note := "See `pkg/alfa/run.go` and `pkg/bravo/stay.go`.\n"
	testutil.WriteFile(t, d, "doc/note.md", note)
	testutil.WriteFile(t, d, "doc/other.md", "Nothing to move here.\n")
	testutil.WriteFile(t, d, "pkg/alfa/run.go", "package alfa\n")
	testutil.WriteFile(t, d, "pkg/bravo/stay.go", "package bravo\n")
	documents := lint.RetargetTree(
		d,
		move_set.New(
			[]*path_move.Move{path_move.New("pkg/alfa", "pkg/charlie")},
			nil,
		),
	)
	assert.Count(t, 1, documents)
	assert.String(t, "doc/note.md", documents[0].Path)
	assert.String(
		t,
		"See `pkg/charlie/run.go` and `pkg/bravo/stay.go`.\n",
		documents[0].Content,
	)
	assert.Count(t, 1, documents[0].Changes)
	assert.String(t, note, system.ReadFile(d, "doc/note.md"))
}
