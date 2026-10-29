package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
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
	w := worker.New(v, time.Hour, logger.New(t.Context()), r)
	w.Start()
	defer w.Close()
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && v.MustStatus().TotalDocuments == 0 {
		time.Sleep(time.Millisecond)
	}

	assert.Integer(t, 1, v.MustStatus().TotalDocuments)
	assert.Count(t, 0, r.Events())
}

func TestTheServiceHandsOutItsEmbedderGuard(t *testing.T) {
	assert.Count(t, 1, newService(t, newEmbedder(t), newReranker(t)).Workers())
}
