package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestARefusedChunkLeavesTheRestOfItsDocumentEmbedded(t *testing.T) {
	body := fixture.Read(systemConstant.LimitPath, "manual.txt")
	assert.Count(t, 3, chunk.Document(body, "manual.txt", mock_reranker.New()))
	m := newEmbedder(t)
	m.Refuse("SYNOPSIS")
	v := newService(t, m, newReranker(t))
	v.PushDocument(t.Context(), "manual", "manual.txt", body, nil)
	assert.Integer(t, 2, v.MustStatus().TotalEmbeddings)
}

func TestARefusedDocumentDoesNotHoldBackThePendingOnes(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("Bravo archive header")
	v := newService(t, m, newReranker(t))
	v.AddCollection("limit", fixture.Path(systemConstant.LimitPath), "**/*")
	_, e := v.IndexCollections("limit")
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, v.MustStatus().PendingEmbeddings)
	_, f := v.Embed(t.Context())
	assert.True(t, validation.Is(f))
	s := v.MustStatus()
	assert.Integer(t, 0, s.PendingEmbeddings)
	assert.Integer(t, 1, s.FailedEmbeddings)
	assert.Integer(t, 3, s.TotalEmbeddings)
}

func TestAChunkOverTheEmbedderWindowNeverReachesTheEmbedder(t *testing.T) {
	m := newEmbedder(t)
	c := newReranker(t)
	c.SetAllowance(300)
	o := service_tester.NewOption(t, m, newReranker(t))
	o.EmbedCounter = func() face.EmbedCounter { return c }
	v := service_tester.New(t, o, memory.New())
	v.PushDocument(
		t.Context(),
		"manual",
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		nil,
	)
	assert.Integer(t, 1, m.Calls())
	s := v.MustStatus()
	assert.Integer(t, 1, s.TotalEmbeddings)
	assert.Integer(t, 2, s.FailedEmbeddings)
	failed, e := v.Failed("")
	assert.FatalOnError(t, e)
	assert.StringContains(t, "over the embedder's window", failed[0].Reason)
	assert.StringContains(t, "tokens against 300", failed[0].Reason)
}

func TestARecordedRefusalIsNotRetried(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("Bravo archive header")
	v := newService(t, m, newReranker(t))
	v.PushDocument(
		t.Context(),
		"dense",
		"dense.md",
		fixture.Read(systemConstant.LimitPath, "dense.md"),
		nil,
	)
	calls := m.Calls()
	_, e := v.Embed(t.Context())
	assert.FatalOnError(t, e)
	assert.Integer(t, calls, m.Calls())
	assert.Integer(t, 1, v.MustStatus().FailedEmbeddings)
}

func TestRechunkLeavesARefusedDocumentAlone(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("SYNOPSIS")
	v := newService(t, m, newReranker(t))
	v.PushDocument(
		t.Context(),
		"manual",
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		nil,
	)
	changed, e := v.Rechunk()
	assert.FatalOnError(t, e)
	assert.Count(t, 0, changed)
	assert.Integer(t, 1, v.MustStatus().FailedEmbeddings)
}

func TestFailedListsTheRefusedChunk(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("SYNOPSIS")
	v := newService(t, m, newReranker(t))
	v.PushDocument(
		t.Context(),
		"manual",
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		nil,
	)
	failed, e := v.Failed("")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, failed)
	assert.String(t, "manual", failed[0].Collection)
	assert.String(t, "manual.txt", failed[0].Path)
	assert.Integer(t, 0, failed[0].Index)
	assert.Integer(t, 1, failed[0].FirstLine)
	assert.Integer(t, 99, failed[0].LastLine)
	assert.StringContains(t, "context length", failed[0].Reason)
	other, f := v.Failed("other")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, other)
}

func TestPruneRemovesAnOldVersionOfAPushedDocument(t *testing.T) {
	v := newService(t, newEmbedder(t), newReranker(t))
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nyankee.\n", nil)
	assert.Integer(t, 1, v.MustStatus().TotalEmbeddings)
	assert.Integer(t, 1, int(v.Prune()))
	assert.Integer(t, 0, int(v.Prune()))
	assert.Integer(t, 1, v.MustStatus().TotalEmbeddings)
	found, _, e := v.GetDocument("notes/alfa.md")
	assert.FatalOnError(t, e)
	assert.StringContains(t, "yankee", found.Body)
}

func TestPruneKeepsTheContentOfARefusedDocument(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("Bravo archive header")
	v := newService(t, m, newReranker(t))
	v.PushDocument(
		t.Context(),
		"dense",
		"dense.md",
		fixture.Read(systemConstant.LimitPath, "dense.md"),
		nil,
	)
	assert.Integer(t, 0, int(v.Prune()))
	assert.Integer(t, 1, v.MustStatus().FailedEmbeddings)
}

func TestDeletingADocumentClearsItsFailures(t *testing.T) {
	m := newEmbedder(t)
	m.Refuse("SYNOPSIS")
	v := newService(t, m, newReranker(t))
	v.PushDocument(
		t.Context(),
		"manual",
		"manual.txt",
		fixture.Read(systemConstant.LimitPath, "manual.txt"),
		nil,
	)
	assert.Integer(t, 1, v.MustStatus().FailedEmbeddings)
	deleted, e := v.DeleteDocument("manual", "manual.txt")
	assert.FatalOnError(t, e)
	assert.True(t, deleted)
	assert.Integer(t, 0, v.MustStatus().FailedEmbeddings)
}
