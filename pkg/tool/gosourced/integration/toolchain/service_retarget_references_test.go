package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"path/filepath"
	"testing"
)

func TestRetargetReferencesFollowsHandMove(t *testing.T) {
	o := module_tester.New(t).HandMoved()
	r, e := o.Reference.RetargetReferences(
		t.Context(),
		o.Directory,
		[]*path_move.Move{
			path_move.New("pkg/alfa/old_name.go", "pkg/alfa/renamed.go"),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "See `pkg/alfa/renamed.go`.\n", o.Read("doc/note.md"))
	assert.Count(t, 1, o.Memory.Retargets)
	assert.String(
		t,
		filepath.Join(o.Directory, "pkg/alfa/renamed.go"),
		o.Memory.Retargets[0].Paths[0].New,
	)
}

func TestRetargetReferencesDryRunWritesNothing(t *testing.T) {
	o := module_tester.New(t).HandMoved()
	r, e := o.Reference.RetargetReferences(
		t.Context(),
		o.Directory,
		[]*path_move.Move{
			path_move.New("pkg/alfa/old_name.go", "pkg/alfa/renamed.go"),
		},
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(t, "See `pkg/alfa/old_name.go`.\n", o.Read("doc/note.md"))
}

func TestRetargetReferencesRefusesPathOutsideModule(t *testing.T) {
	o := module_tester.New(t).HandMoved()
	r, e := o.Reference.RetargetReferences(
		t.Context(),
		o.Directory,
		[]*path_move.Move{path_move.New("../elsewhere.md", "doc/here.md")},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"../elsewhere.md is outside the module",
	)
}

func TestRetargetReferencesSaysWhenNothingCitesMove(t *testing.T) {
	o := module_tester.New(t).HandMoved()
	r, e := o.Reference.RetargetReferences(
		t.Context(),
		o.Directory,
		[]*path_move.Move{
			path_move.New("pkg/alfa/pair.go", "pkg/alfa/couple.go"),
		},
		false,
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"nothing references the moved paths",
		rewrite_tester.JoinedConcerns(r),
	)
}
