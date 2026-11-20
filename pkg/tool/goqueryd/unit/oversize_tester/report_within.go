package oversize_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
)

func (o *Tester) ReportWithin(
	collection string,
	tokenLimit int,
	embedTokenLimit int,
) *oversize.Report {
	o.t.Helper()
	q := service_tester.New(o.t)
	q.Reranker.SetTokenLimit(tokenLimit)
	q.Counter.SetTokenLimit(embedTokenLimit)
	w := store_tester.New(o.t)
	s := w.OpenFileStore(q.Path)
	d := o.t.TempDir()
	w.WriteFixture(d, "alfa.md", "one two three\n")
	w.WriteFixture(
		d,
		"bravo.md",
		"one two\nthree four\nfive six\nseven eight nine ten\n",
	)
	w.WriteFixture(d, "charlie.md", "one two three four\nfive six seven\n")
	s.AddCollection("test", d, constant.DefaultGlob)
	s.Index("test")
	result, e := q.Service().Oversize(collection)
	assert.FatalOnError(o.t, e)

	return result
}
