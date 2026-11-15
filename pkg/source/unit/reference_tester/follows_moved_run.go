package reference_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func FollowsMovedRun(
	t *testing.T,
	bravo string,
	charlie string,
) {
	t.Helper()
	directory := t.TempDir()
	system_tester.Write(
		t,
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\nfunc (Inner) Run() {}\n\ntype Outer struct {\n\tInner\n}\n",
	)
	system_tester.Write(t, directory, "bravo/bravo.go", bravo)
	system_tester.Write(t, directory, "charlie/call.go", charlie)
	testutil.WriteModFile(t, directory, "example")
	store := t.TempDir()
	assert.Strings(
		t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			Target(t, directory, "Outer", "Run"),
		),
	)
	system_tester.Write(
		t,
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\ntype Outer struct {\n\tInner\n}\n\nfunc (Outer) Run() {}\n",
	)
	assert.Strings(
		t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			Target(t, directory, "Outer", "Run"),
		),
	)
}
