package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConcurrentProjectLookupsShareTheCache(t *testing.T) {
	var hits atomic.Int64
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					hits.Add(1)
					web.ObjectHeader(w)
					writer.Print(
						w,
						`{"id": 7, "path": "seven", "namespace": {"path": "group"}}`,
					)
				},
			)
		},
	)
	var g sync.WaitGroup

	for range 8 {
		g.Go(
			func() {
				_, e := r.Client.Project(7)
				assert.FatalOnError(t, e)
			},
		)
	}

	g.Wait()
	hits.Store(0)
	p, e := r.Client.Project(7)
	assert.FatalOnError(t, e)
	assert.String(t, "seven", p.Name)
	assert.Integer(t, 0, int(hits.Load()))
}
