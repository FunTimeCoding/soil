package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	"net/http"
	"testing"
	"time"
)

func TestARefusedSearchEmbeddingIsReportedOnceBySurface(t *testing.T) {
	m := newEmbedder(t)
	m.FailEmbedding(errors.New("embedder exploded"))
	r := memory.New()
	v := service_tester.New(
		t,
		service_tester.NewOption(t, m, newReranker(t)),
		r,
	)
	c := mountedClient(t, v, r)
	answer, e := c.GetSearchWithResponse(
		t.Context(),
		&client.GetSearchParams{Query: "zulu"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, answer.StatusCode())
	assert.Count(t, 1, r.Events())
}

func TestTheWorkerReportsAnUnexpectedEmbeddingFailure(t *testing.T) {
	m := newEmbedder(t)
	m.FailEmbedding(errors.New("embedder exploded"))
	r := memory.New()
	l := logger.New(context.Background())
	v := service_tester.New(
		t,
		service_tester.NewOption(t, m, newReranker(t)),
		r,
	)
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	assert.Count(t, 0, r.Events())
	w := worker.New(v, time.Hour, l, r)
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && len(r.Events()) == 0 {
		time.Sleep(time.Millisecond)
	}

	assert.Count(t, 1, r.Events())
	assert.StringContains(t, "embedder exploded", r.Events()[0].Error.Error())
}

func TestTheWorkerLeavesAQuietEmbeddingRefusalOutOfSentry(t *testing.T) {
	m := newEmbedder(t)
	m.FailEmbedding(validation.New("input too long"))
	r := memory.New()
	l := logger.New(context.Background())
	v := service_tester.New(
		t,
		service_tester.NewOption(t, m, newReranker(t)),
		r,
	)
	directory := t.TempDir()
	store_tester.WriteFixture(t, directory, "alfa.md", "# Alfa\n\nzulu.\n")
	v.AddCollection("notes", directory, constant.DefaultGlob)
	w := worker.New(v, 5*time.Millisecond, l, r)
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && v.MustStatus().FailedEmbeddings == 0 {
		time.Sleep(time.Millisecond)
	}

	assert.Integer(t, 1, v.MustStatus().FailedEmbeddings)
	assert.Count(t, 0, r.Events())
}
