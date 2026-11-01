package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestDiffSymbolsComparesHeadWithTheIndex(t *testing.T) {
	r := diffRepository(t)
	v := diffService(r.Clone)
	result, e := v.DiffSymbols("echo", "", "")
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"unchanged . alfa Charlie", "body . alfa Bravo"},
		reportLines(result),
	)
}

func TestDiffSymbolsComparesTwoCommits(t *testing.T) {
	r := diffRepository(t)
	r.Commit("bravo")
	r.Write("alfa.go", "package alfa\n\nfunc Charlie() int {\n\treturn 3\n}\n")
	r.Commit("charlie")
	v := diffService(r.Clone)
	result, e := v.DiffSymbols("echo", "HEAD~2", "HEAD")
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"unchanged . alfa Charlie", "removed . alfa Bravo"},
		reportLines(result),
	)
}

func TestDiffSymbolsRefusesAnUnknownModuleOrRevision(t *testing.T) {
	r := diffRepository(t)
	v := diffService(r.Clone)
	_, e := v.DiffSymbols("foxtrot", "", "")
	assert.True(t, not_found.Is(e))
	_, f := v.DiffSymbols("echo", "nowhere", "")
	assert.True(t, not_found.Is(f))
}
