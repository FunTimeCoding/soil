package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
	"testing"
	"time"
)

func TestTheWorkerIndexesWhileTheEmbedderIsDown(t *testing.T) {
	m := newEmbedder(t)
	m.Fail(refusedConnection())
	v := newService(t, m, newReranker(t))
	d := t.TempDir()
	store_tester.WriteFixture(t, d, "alfa.md", "# Alfa\n\nzulu.\n")
	v.AddCollection("test", d, constant.DefaultGlob)
	r := memory.New()
	w := worker.New(
		v.Maintain,
		time.Hour,
		logger.New(t.Context()),
		prometheus.NewRegistry(),
		r,
	)
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && v.MustStatus().TotalDocuments == 0 {
		time.Sleep(time.Millisecond)
	}

	assert.Integer(t, 1, v.MustStatus().TotalDocuments)
	assert.Count(t, 0, r.Events())
}

func TestTheWorkerCountsWhatItPrunes(t *testing.T) {
	v := newService(t, newEmbedder(t), newReranker(t))
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nzulu.\n", nil)
	v.PushDocument(t.Context(), "notes", "alfa.md", "# Alfa\n\nyankee.\n", nil)
	g := prometheus.NewRegistry()
	w := worker.New(
		v.Maintain,
		time.Hour,
		logger.New(t.Context()),
		g,
		memory.New(),
	)
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && prunedContents(g) == 0 {
		time.Sleep(time.Millisecond)
	}

	assert.Float(t, 1, prunedContents(g))
}

func TestTheWorkerCloseWaitsForTheRunningPoll(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	closed := make(chan struct{})
	var completed atomic.Int64
	w := worker.New(
		func(context.Context) (int64, error) {
			close(entered)
			<-release
			completed.Add(1)

			return 0, nil
		},
		time.Hour,
		logger.New(t.Context()),
		prometheus.NewRegistry(),
		memory.New(),
	)
	w.Start()
	<-entered
	go func() {
		w.Close()
		close(closed)
	}()

	select {
	case <-closed:
		t.Fatal("close returned while the poll was running")
	case <-time.After(50 * time.Millisecond):
	}

	close(release)

	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("close did not return after the poll finished")
	}

	assert.Integer(t, 1, int(completed.Load()))
}
