package reference_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
)

func (o *Tester) FollowsMovedRun(
	bravo string,
	charlie string,
) {
	o.t.Helper()
	directory := o.t.TempDir()
	system_tester.New(o.t).Write(
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\nfunc (Inner) Run() {}\n\ntype Outer struct {\n\tInner\n}\n",
	)
	system_tester.New(o.t).Write(directory, "bravo/bravo.go", bravo)
	system_tester.New(o.t).Write(directory, "charlie/call.go", charlie)
	testutil.WriteModFile(o.t, directory, "example")
	store := o.t.TempDir()
	assert.Strings(
		o.t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			o.Target(directory,"Outer", "Run"),
		),
	)
	system_tester.New(o.t).Write(
		directory,
		"alfa/outer.go",
		"package alfa\n\ntype Inner struct{}\n\ntype Outer struct {\n\tInner\n}\n\nfunc (Outer) Run() {}\n",
	)
	assert.Strings(
		o.t,
		[]string{"example/charlie"},
		index.New(store, directory).References().Referencing(
			o.Target(directory,"Outer", "Run"),
		),
	)
}
