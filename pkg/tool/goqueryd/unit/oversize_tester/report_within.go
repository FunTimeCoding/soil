package oversize_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func ReportWithin(
	t *testing.T,
	collection string,
	tokenLimit int,
	embedTokenLimit int,
) *oversize.Report {
	t.Helper()
	q := service_tester.New(t)
	q.Reranker.SetTokenLimit(tokenLimit)
	q.Counter.SetTokenLimit(embedTokenLimit)
	s := store_tester.OpenFileStore(t, q.Path)
	d := t.TempDir()
	store_tester.WriteFixture(t, d, "alfa.md", "one two three\n")
	store_tester.WriteFixture(
		t,
		d,
		"bravo.md",
		"one two\nthree four\nfive six\nseven eight nine ten\n",
	)
	store_tester.WriteFixture(
		t,
		d,
		"charlie.md",
		"one two three four\nfive six seven\n",
	)
	s.AddCollection("test", d, constant.DefaultGlob)
	s.Index("test")
	result, e := q.Service().Oversize(collection)
	assert.FatalOnError(t, e)

	return result
}
