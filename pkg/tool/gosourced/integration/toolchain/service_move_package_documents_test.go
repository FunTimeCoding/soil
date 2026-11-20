package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestMovePackageRetargetsDocuments(t *testing.T) {
	o := module_tester.New(t).PackageMoveDocument()
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		"Store lives in `pkg/store/store.go`, built by `go:pkg/store.New`, called from `pkg/caller/run.go`.\n",
		o.Read("doc/note.md"),
	)
}

func TestMovePackageDryRunLeavesDocuments(t *testing.T) {
	o := module_tester.New(t).PackageMoveDocument()
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
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
		o.Read("doc/note.md"),
	)
}

func TestMovePackageKeepsDocumentEditedMeanwhile(t *testing.T) {
	o := module_tester.New(t).PackageMoveDocument()
	o.Commit.BeforeCommit(
		func() {
			system_tester.New(t).Write(
				o.Directory,
				"doc/note.md",
				"Edited meanwhile.\n",
			)
		},
	)
	r, e := o.Packages.MovePackage(
		t.Context(),
		o.Directory,
		"example/pkg/outer/store",
		"example/pkg/store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.StringContains(
		t,
		"doc/note.md: changed while the call ran - ran again",
		strings.Join(service_tester.ConcernText(r), stringsConstant.Unix),
	)
	assert.String(t, "Edited meanwhile.\n", o.Read("doc/note.md"))
}
