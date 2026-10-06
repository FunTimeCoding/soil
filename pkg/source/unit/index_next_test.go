package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/source/unit/reference_tester"
	"os"
	"testing"
)

func TestNextAnswersLikeAFreshWorkspace(t *testing.T) {
	directory := reference_tester.Module(t)
	store := t.TempDir()
	start := reference_tester.Target(t, directory, "Server", "Start")
	previous := index.New(store, directory)
	before := previous.References().Locations(start)
	taken := snapshot.Take(directory)
	testutil.WriteFile(
		t,
		directory,
		"bravo/use.go",
		"package bravo\n\nimport \"example/alfa\"\n\nfunc Use() int {\n\ts := alfa.NewServer()\n\ts.Start()\n\ts.Start()\n\tb := &alfa.Box[int]{}\n\n\treturn s.Port + b.Get()\n}\n",
	)
	next := index.Next(
		previous,
		taken.Diff(snapshot.Take(directory)),
	).References().Locations(start)
	assert.Strings(
		t,
		index.New(t.TempDir(), directory).References().Locations(start),
		next,
	)
	assert.Integer(t, len(before)+1, len(next))
}

func TestNextCarriesFactsOfUnchangedPackages(t *testing.T) {
	directory := kindModule(t)
	store := t.TempDir()
	var extracted, external int
	alfa := countingKind("alfa", &extracted, &external)
	previous := index.New(store, directory, alfa)
	index.Facts[*string](previous, alfa)
	assert.Integer(t, 2, extracted)
	assert.FatalOnError(t, os.RemoveAll(store))
	taken := snapshot.Take(directory)
	testutil.WriteFile(
		t,
		directory,
		"two/two.go",
		"package two\n\nimport \"example/one\"\n\nfunc Two() string {\n\treturn one.One() + \"!\"\n}\n",
	)
	next := index.Next(previous, taken.Diff(snapshot.Take(directory)))
	facts := index.Facts[*string](next, alfa)
	assert.Integer(t, 2, len(facts))
	assert.Integer(t, 3, extracted)
}
