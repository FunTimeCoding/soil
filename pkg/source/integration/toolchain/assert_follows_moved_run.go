package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/reference_tester"
	"testing"
)

func assertFollowsMovedRun(
	t *testing.T,
	bravo string,
	charlie string,
) {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(
		t,
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\nfunc (Inner) Run() {}\n\ntype Outer struct {\n\tInner\n}\n",
	)
	testutil.WriteFile(t, directory, "bravo/bravo.go", bravo)
	testutil.WriteFile(t, directory, "charlie/call.go", charlie)
	testutil.WriteModFile(t, directory, "example")
	store := t.TempDir()
	assert.Strings(
		t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			reference_tester.Target(t, directory, "Outer", "Run"),
		),
	)
	testutil.WriteFile(
		t,
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\ntype Outer struct {\n\tInner\n}\n\nfunc (Outer) Run() {}\n",
	)
	assert.Strings(
		t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			reference_tester.Target(t, directory, "Outer", "Run"),
		),
	)
}
