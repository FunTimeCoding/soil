package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRetargetReferencesFollowsHandMove(t *testing.T) {
	d := handMovedFixture(t)
	s, m := service_tester.WithMemory(t)
	r, e := s.Reference.RetargetReferences(
		t.Context(),
		d,
		[]*path_move.Move{
			path_move.New("pkg/alfa/old_name.go", "pkg/alfa/renamed.go"),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"See `pkg/alfa/renamed.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
	assert.Count(t, 1, m.Retargets)
	assert.String(
		t,
		filepath.Join(d, "pkg/alfa/renamed.go"),
		m.Retargets[0].Paths[0].New,
	)
}

func TestRetargetReferencesDryRunWritesNothing(t *testing.T) {
	d := handMovedFixture(t)
	r, e := service_tester.NewService(t).Reference.RetargetReferences(
		t.Context(),
		d,
		[]*path_move.Move{
			path_move.New("pkg/alfa/old_name.go", "pkg/alfa/renamed.go"),
		},
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"See `pkg/alfa/old_name.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestRetargetReferencesRefusesPathOutsideModule(t *testing.T) {
	d := handMovedFixture(t)
	r, e := service_tester.NewService(t).Reference.RetargetReferences(
		t.Context(),
		d,
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
	d := handMovedFixture(t)
	r, e := service_tester.NewService(t).Reference.RetargetReferences(
		t.Context(),
		d,
		[]*path_move.Move{
			path_move.New("pkg/alfa/pair.go", "pkg/alfa/couple.go"),
		},
		false,
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"nothing references the moved paths",
		joinedConcerns(r),
	)
}
