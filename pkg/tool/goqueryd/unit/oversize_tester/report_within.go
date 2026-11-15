package oversize_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/oversize"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func ReportWithin(
	t *testing.T,
	collection string,
	allowance int,
	embedAllowance int,
) *oversize.Report {
	t.Helper()
	a := service_tester.NewReranker(t)
	a.SetAllowance(allowance)
	c := service_tester.NewReranker(t)
	c.SetAllowance(embedAllowance)
	o := service_tester.NewOption(t, service_tester.NewEmbedder(t), a)
	o.EmbedCounter = func() face.EmbedCounter { return c }
	s := store_tester.OpenFileStore(t, o.LitePath)
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
	result, e := service_tester.New(t, o, memory.New()).Oversize(collection)
	assert.FatalOnError(t, e)

	return result
}
