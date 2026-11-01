package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestMovePackageRetargetsDocuments(t *testing.T) {
	d := packageMoveDocumentFixture(t)
	r, e := service_tester.NewService().MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Store lives in `pkg/store/store.go`, built by `go:pkg/store.New`, called from `pkg/caller/run.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMovePackageDryRunLeavesDocuments(t *testing.T) {
	d := packageMoveDocumentFixture(t)
	r, e := service_tester.NewService().MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"pkg/outer/store/store.go → pkg/store/store.go",
		strings.Join(service_tester.ConcernText(r), stringsConstant.Unix),
	)
	assert.String(
		t,
		"Store lives in `pkg/outer/store/store.go`, built by `go:pkg/outer/store.New`, called from `pkg/caller/run.go`.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}

func TestMovePackageRefusesWhenADocumentChanges(t *testing.T) {
	d := packageMoveDocumentFixture(t)
	s := service_tester.NewService()
	s.BeforeCommit(
		func() {
			testutil.WriteFile(t, d, "doc/note.md", "Edited meanwhile.\n")
		},
	)
	r, e := s.MovePackage(
		d,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "changed since it was read")
	assert.String(
		t,
		"Edited meanwhile.\n",
		service_tester.ReadFixtureFile(t, d, "doc/note.md"),
	)
}
