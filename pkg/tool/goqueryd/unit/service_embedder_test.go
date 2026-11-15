package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestPushIsEmbeddedWhileEmbedderAnswers(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	assert.String(
		t,
		"",
		v.MustPushDocument(
			t.Context(),
			"notes",
			"alfa.md",
			"# Alfa\n\nzulu.\n",
			nil,
		),
	)
	assert.Integer(t, 0, v.MustStatus().PendingEmbeddings)
}

func TestPushWhileEmbedderIsDownIsStoredWithItsMetadata(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	warning := v.MustPushDocument(
		t.Context(),
		"notes",
		"alfa.md",
		"# Alfa\n\nzulu.\n",
		map[string][]string{"scope": {"alfa"}},
	)
	assert.StringContains(t, "stored; embedding pending", warning)
	assert.StringContains(t, "embedder unreachable since", warning)
	assert.Integer(t, 1, v.MustStatus().PendingEmbeddings)
	listed, e := v.ListDocuments(
		"notes",
		map[string]string{"scope": "alfa"},
		0,
		0,
		false,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, listed.Results)
}

func TestHybridSearchFailsAtOnceWhileEmbedderIsDown(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	_, e := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.True(t, unavailable.Is(e))
	_, f := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "embedder unreachable since", f.Error())
	assert.Integer(t, 1, m.Calls())
}

func TestKeywordSearchNeedsNoEmbedder(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	v.MustPushDocument(
		t.Context(),
		"notes",
		"alfa.md",
		"# Alfa\n\nzulu.\n",
		nil,
	)
	o := search_option.New("zulu", 10)
	o.Mode = "keyword"
	found, e := v.Search(t.Context(), o)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found.Results)
}

func TestPendingDocumentsWaitWhileEmbedderIsDown(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	m.Fail(outage_tester.RefusedDial())
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	v.MustPushDocument(
		t.Context(),
		"notes",
		"alfa.md",
		"# Alfa\n\nzulu.\n",
		nil,
	)
	_, e := v.Embed(t.Context())
	assert.True(t, unavailable.Is(e))
	assert.Integer(t, 1, v.MustStatus().PendingEmbeddings)
}
