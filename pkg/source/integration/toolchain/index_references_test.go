package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/reference_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func TestReferencingCoversEveryKindOfUnit(t *testing.T) {
	directory := reference_tester.New(t).Module()
	i := index.New(t.TempDir(), directory).References()
	assert.Strings(
		t,
		[]string{"example/bravo", "example/bravo_test", "example/charlie/unit"},
		i.Referencing(
			reference_tester.New(t).Target(directory, "NewServer", ""),
		),
	)
}

func TestReferencingFindsMembers(t *testing.T) {
	directory := reference_tester.New(t).Module()
	i := index.New(t.TempDir(), directory).References()
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(
			reference_tester.New(t).Target(directory, "Server", "Start"),
		),
	)
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(
			reference_tester.New(t).Target(directory, "Server", "Port"),
		),
	)
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(reference_tester.New(t).Target(directory, "Box", "Get")),
	)
}

func TestReferencingFollowsNewPackage(t *testing.T) {
	directory := reference_tester.New(t).Module()
	store := t.TempDir()
	index.New(store, directory).References()
	system_tester.New(t).Write(
		directory,
		"delta/late.go",
		"package delta\n\nimport \"example/alfa\"\n\nfunc Late() {\n\t_ = alfa.NewServer()\n}\n",
	)
	assert.Strings(
		t,
		[]string{
			"example/bravo",
			"example/bravo_test",
			"example/charlie/unit",
			"example/delta",
		},
		index.New(store, directory).References().Referencing(
			reference_tester.New(t).Target(directory, "NewServer", ""),
		),
	)
}

func TestReferencingFollowsMovedMethodThroughField(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\ntype Holder struct {\n\tOuter *alfa.Outer\n}\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call(h *bravo.Holder) {\n\th.Outer.Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughResult(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\nfunc Make() alfa.Outer {\n\treturn alfa.Outer{}\n}\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call() {\n\tbravo.Make().Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughEmbedding(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\ntype Holder struct {\n\talfa.Outer\n}\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call(h bravo.Holder) {\n\th.Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughTypeArgument(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport (\n\t\"example/alfa\"\n\t\"sync/atomic\"\n)\n\nvar Default atomic.Pointer[alfa.Outer]\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call() {\n\tbravo.Default.Load().Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughMap(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\nvar Registry map[string]alfa.Outer\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call() {\n\tbravo.Registry[\"x\"].Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughAlias(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\ntype Outer = alfa.Outer\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call(o bravo.Outer) {\n\to.Run()\n}\n",
	)
}

func TestReferencingFollowsMovedMethodThroughInterface(t *testing.T) {
	reference_tester.New(t).FollowsMovedRun(
		"package bravo\n\nimport \"example/alfa\"\n\nfunc Source() interface{ Outer() alfa.Outer } {\n\treturn nil\n}\n",
		"package charlie\n\nimport \"example/bravo\"\n\nfunc Call() {\n\tbravo.Source().Outer().Run()\n}\n",
	)
}

func TestWorkspaceSharesOneRefresh(t *testing.T) {
	directory := reference_tester.New(t).Module()
	w := index.New(t.TempDir(), directory)
	first := w.References()
	assert.True(t, first == w.References())
}
