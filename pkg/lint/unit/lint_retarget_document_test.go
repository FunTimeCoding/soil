package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"testing"
)

func TestRetargetRepositoryPathKeepsItsLine(t *testing.T) {
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/alfa/bravo.go"),
		"doc/note.md",
		"Read `pkg/alfa/bravo.go:12` first.\n",
		[]*path_move.Move{path_move.New("pkg/alfa", "pkg/charlie")},
	)
	assert.String(t, "Read `pkg/charlie/bravo.go:12` first.\n", after)
	assert.Count(t, 1, changes)
}

func TestRetargetRenamedFile(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/alfa/old_name.go"),
		"doc/note.md",
		"See `pkg/alfa/old_name.go`.\n",
		[]*path_move.Move{
			path_move.New("pkg/alfa/old_name.go", "pkg/alfa/new_name.go"),
		},
	)
	assert.String(t, "See `pkg/alfa/new_name.go`.\n", after)
}

func TestRetargetSiblingDirectory(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("../soil/pkg/alfa"),
		"doc/note.md",
		"Lives in `../soil/pkg/alfa/`.\n",
		[]*path_move.Move{
			path_move.New("../soil/pkg/alfa", "../soil/pkg/bravo"),
		},
	)
	assert.String(t, "Lives in `../soil/pkg/bravo/`.\n", after)
}

func TestRetargetSymbolPackage(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver(),
		"doc/note.md",
		"Grep `go:pkg/alfa.Value` for it.\n",
		[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
	)
	assert.String(t, "Grep `go:pkg/bravo.Value` for it.\n", after)
}

func TestRetargetBaseMovesWithItsShortPaths(t *testing.T) {
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver(
			"pkg/alfa",
			"pkg/alfa/service",
			"pkg/alfa/service/run.go",
		),
		"doc/note.md",
		"---\nbase: pkg/alfa\n---\n\nSee `service/run.go`.\n",
		[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
	)
	assert.String(
		t,
		"---\nbase: pkg/bravo\n---\n\nSee `service/run.go`.\n",
		after,
	)
	assert.Count(t, 0, changes)
}

func TestRetargetShortPathLeavingItsBase(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver(
			"pkg/alfa",
			"pkg/alfa/service",
			"pkg/alfa/service/run.go",
		),
		"doc/note.md",
		"---\nbase: pkg/alfa\n---\n\nSee `service/run.go`.\n",
		[]*path_move.Move{
			path_move.New("pkg/alfa/service", "pkg/bravo/service"),
		},
	)
	assert.String(
		t,
		"---\nbase: pkg/alfa\n---\n\nSee `pkg/bravo/service/run.go`.\n",
		after,
	)
}

func TestRetargetRelativeLink(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/alfa/run.go"),
		"doc/ai/note.md",
		"The [runner](../../pkg/alfa/run.go) starts it.\n",
		[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
	)
	assert.String(t, "The [runner](../../pkg/bravo/run.go) starts it.\n", after)
}

func TestRetargetBracedDirectory(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("doc/alfa/bravo.md", "doc/alfa/charlie.md"),
		"doc/note.md",
		"Read `doc/alfa/{bravo,charlie}.md`.\n",
		[]*path_move.Move{path_move.New("doc/alfa", "doc/delta")},
	)
	assert.String(t, "Read `doc/delta/{bravo,charlie}.md`.\n", after)
}

func TestRetargetBracedSplitIsLeftForHand(t *testing.T) {
	content := "Read `doc/alfa/{bravo,charlie}.md`.\n"
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver("doc/alfa/bravo.md", "doc/alfa/charlie.md"),
		"doc/note.md",
		content,
		[]*path_move.Move{
			path_move.New("doc/alfa/bravo.md", "doc/echo/bravo.md"),
		},
	)
	assert.String(t, content, after)
	assert.Count(t, 1, changes)
	assert.String(t, "braced span needs a hand edit", changes[0].Reason)
}

func TestRetargetKeepsTheFragment(t *testing.T) {
	after, _ := lint.RetargetDocument(
		pointer_tester.Resolver("doc/alfa.md"),
		"doc/note.md",
		"See `doc/alfa.md#usage`.\n",
		[]*path_move.Move{path_move.New("doc/alfa.md", "doc/bravo.md")},
	)
	assert.String(t, "See `doc/bravo.md#usage`.\n", after)
}

func TestRetargetLeavesProseAndUnrelatedPaths(t *testing.T) {
	content := "The pkg/alfa tree and `pkg/alfabet/run.go` stay.\n"
	after, changes := lint.RetargetDocument(
		pointer_tester.Resolver("pkg/alfabet/run.go"),
		"doc/note.md",
		content,
		[]*path_move.Move{path_move.New("pkg/alfa", "pkg/bravo")},
	)
	assert.String(t, content, after)
	assert.Count(t, 0, changes)
}
