package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/reference_tester"
	"testing"
)

func TestReferencingCoversEveryKindOfUnit(t *testing.T) {
	directory := reference_tester.Module(t)
	i := index.New(t.TempDir(), directory).References()
	assert.Strings(
		t,
		[]string{"example/bravo", "example/bravo_test", "example/charlie/unit"},
		i.Referencing(reference_tester.Target(t, directory, "NewServer", "")),
	)
}

func TestReferencingFindsMembers(t *testing.T) {
	directory := reference_tester.Module(t)
	i := index.New(t.TempDir(), directory).References()
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(reference_tester.Target(t, directory, "Server", "Start")),
	)
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(reference_tester.Target(t, directory, "Server", "Port")),
	)
	assert.Strings(
		t,
		[]string{"example/bravo"},
		i.Referencing(reference_tester.Target(t, directory, "Box", "Get")),
	)
}

func TestReferencingFollowsNewPackage(t *testing.T) {
	directory := reference_tester.Module(t)
	store := t.TempDir()
	index.New(store, directory).References()
	testutil.WriteFile(
		t,
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
			reference_tester.Target(t, directory, "NewServer", ""),
		),
	)
}

func TestWorkspaceSharesOneRefresh(t *testing.T) {
	directory := reference_tester.Module(t)
	w := index.New(t.TempDir(), directory)
	first := w.References()
	assert.True(t, first == w.References())
}
