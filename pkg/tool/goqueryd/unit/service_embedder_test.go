package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"testing"
)

func TestAPushIsEmbeddedWhileTheEmbedderAnswers(t *testing.T) {
	m := newEmbedder(t)
	v := newService(t, m, newReranker(t))
	assert.String(
		t,
		"",
		v.PushDocument(
			t.Context(),
			"notes",
			"alfa.md",
			"# Alfa\n\nzulu.\n",
			nil,
		),
	)
	assert.Integer(t, 0, v.MustStatus().PendingEmbeddings)
}

func TestAPushWhileTheEmbedderIsDownIsStoredWithItsMetadata(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	warning := v.PushDocument(
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

func TestAHybridSearchFailsAtOnceWhileTheEmbedderIsDown(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	_, e := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.True(t, unavailable.Is(e))
	_, f := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "embedder unreachable since", f.Error())
	assert.Integer(t, 1, m.Calls())
}

func TestAKeywordSearchNeedsNoEmbedder(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	o := search_option.New("zulu", 10)
	o.Mode = "keyword"
	found, e := v.Search(t.Context(), o)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, found.Results)
}

func TestPendingDocumentsWaitWhileTheEmbedderIsDown(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	_, e := v.Embed(t.Context())
	assert.True(t, unavailable.Is(e))
	assert.Integer(t, 1, v.MustStatus().PendingEmbeddings)
}
