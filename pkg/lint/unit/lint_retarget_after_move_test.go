package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"testing"
)

func TestRetargetAfterTheMoveHappened(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/bravo/run.go"),
		"doc/note.md",
		"Read `pkg/alfa/run.go` first.\n",
		pathMoves(path_move.New("pkg/alfa", "pkg/bravo")),
	)
	assert.String(t, "Read `pkg/bravo/run.go` first.\n", after)
}

func TestRetargetMovedBaseAfterTheMove(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver(
			"pkg/bravo",
			"pkg/bravo/service",
			"pkg/bravo/service/run.go",
		),
		"doc/note.md",
		"---\nbase: pkg/alfa\n---\n\nSee `service/run.go`.\n",
		pathMoves(path_move.New("pkg/alfa", "pkg/bravo")),
	)
	assert.String(
		t,
		"---\nbase: pkg/bravo\n---\n\nSee `service/run.go`.\n",
		after,
	)
}

func TestRetargetLeavesADeadPathThatNoMoveExplains(t *testing.T) {
	content := "Read `pkg/charlie/run.go` first.\n"
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/bravo/run.go"),
		"doc/note.md",
		content,
		pathMoves(path_move.New("pkg/alfa", "pkg/bravo")),
	)
	assert.String(t, content, after)
	assert.Count(t, 0, changes)
}
