package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func oversizeReport(
	t *testing.T,
	collection string,
) *oversize.Report {
	t.Helper()
	s := store_tester.OpenTestStore(t)
	t.Cleanup(s.Close)
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
	a := mock_reranker.New()
	a.SetAllowance(5)
	result, e := service.New(s, ollama.New(), a).Oversize(collection)
	assert.FatalOnError(t, e)

	return result
}
