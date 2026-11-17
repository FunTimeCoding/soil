package repository

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/report_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/revision_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestDiffSymbolsComparesHeadWithIndex(t *testing.T) {
	r := revision_tester.DiffRepository(t)
	v := service_tester.New(
		t,
		service_tester.WithInventory(service_tester.Inventory(r.Clone)),
	)
	result, e := v.Diff.DiffSymbols("echo", "", "")
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"unchanged . alfa Charlie", "body . alfa Bravo"},
		report_tester.Lines(result),
	)
}

func TestDiffSymbolsComparesTwoCommits(t *testing.T) {
	r := revision_tester.DiffRepository(t)
	r.Commit("bravo")
	r.Write("alfa.go", "package alfa\n\nfunc Charlie() int {\n\treturn 3\n}\n")
	r.Commit("charlie")
	v := service_tester.New(
		t,
		service_tester.WithInventory(service_tester.Inventory(r.Clone)),
	)
	result, e := v.Diff.DiffSymbols("echo", "HEAD~2", "HEAD")
	assert.FatalOnError(t, e)
	assert.Strings(
		t,
		[]string{"unchanged . alfa Charlie", "removed . alfa Bravo"},
		report_tester.Lines(result),
	)
}

func TestDiffSymbolsRefusesUnknownModuleOrRevision(t *testing.T) {
	r := revision_tester.DiffRepository(t)
	v := service_tester.New(
		t,
		service_tester.WithInventory(service_tester.Inventory(r.Clone)),
	)
	_, e := v.Diff.DiffSymbols("foxtrot", "", "")
	assert.True(t, not_found.Is(e))
	_, f := v.Diff.DiffSymbols("echo", "nowhere", "")
	assert.True(t, not_found.Is(f))
}
