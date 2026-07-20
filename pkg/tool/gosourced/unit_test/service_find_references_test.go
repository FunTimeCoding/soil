package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"testing"
)

func TestFindReferences(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FindReferences(d, "example/pkg/target", "Used", "")
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, references)
	assert.String(t, "Used", references.Symbol)
	assert.Integer(t, 1, references.Total)
	assert.Integer(t, 1, len(references.Locations))
	assert.String(t, "pkg/caller/run.go", references.Locations[0].File)
	assert.Integer(t, 8, references.Locations[0].Line)
	assert.String(t, "example/pkg/caller", references.Locations[0].Package)
}

func TestFindReferencesMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FindReferences(
		d,
		"example/pkg/target",
		"Ping",
		"Thing",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, references)
	assert.Integer(t, 1, references.Total)
	assert.String(t, "pkg/caller/run.go", references.Locations[0].File)
}

func TestFindReferencesUnknownSymbol(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FindReferences(
		d,
		"example/pkg/target",
		"Missing",
		"",
	)
	assert.FatalOnError(t, e)
	assert.True(t, references == nil)
	testutil.AssertBlockedContains(t, r, "Missing")
}

func TestFileReferences(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FileReferences(
		d,
		"example/pkg/target",
		"pkg/target/constant.go",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, references)
	assert.String(t, "pkg/target/constant.go", references.File)
	assert.Integer(t, 2, len(references.Symbols))
	assert.String(t, "Used", references.Symbols[0].Symbol)
	assert.Integer(t, 1, references.Symbols[0].Total)
	assert.String(t, "Unused", references.Symbols[1].Symbol)
	assert.Integer(t, 0, references.Symbols[1].Total)
}

func TestFileReferencesMethods(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FileReferences(
		d,
		"example/pkg/target",
		"pkg/target/thing.go",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.NotNil(t, references)
	assert.Integer(t, 2, len(references.Symbols))
	assert.String(t, "Thing", references.Symbols[0].Symbol)
	assert.Integer(t, 1, references.Symbols[0].Total)
	assert.String(t, "Thing.Ping", references.Symbols[1].Symbol)
	assert.Integer(t, 1, references.Symbols[1].Total)
}

func TestFileReferencesUnknownFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		serviceTestdata("find-references/src"),
	)
	s := testService()
	r, references, e := s.FileReferences(
		d,
		"example/pkg/target",
		"pkg/target/missing.go",
	)
	assert.FatalOnError(t, e)
	assert.True(t, references == nil)
	testutil.AssertBlockedContains(t, r, "missing.go")
}

func TestPaginateReferences(t *testing.T) {
	references := result.NewReferences(
		"Used",
		[]*result.Location{
			result.NewLocation("a.go", 1, ""),
			result.NewLocation("b.go", 2, ""),
			result.NewLocation("c.go", 3, ""),
			result.NewLocation("d.go", 4, ""),
		},
	)
	references.Paginate(2, 1)
	assert.Integer(t, 4, references.Total)
	assert.Integer(t, 2, len(references.Locations))
	assert.String(t, "b.go", references.Locations[0].File)
	assert.String(t, "c.go", references.Locations[1].File)
	assert.Integer(t, 1, references.More)
}

func TestPaginateReferencesOffsetBeyond(t *testing.T) {
	references := result.NewReferences(
		"Used",
		[]*result.Location{result.NewLocation("a.go", 1, "")},
	)
	references.Paginate(25, 5)
	assert.Integer(t, 0, len(references.Locations))
	assert.Integer(t, 0, references.More)
}
