package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"net/http"
	"testing"
	"time"
)

func TestACancelledCallStopsWaitingOnGitLab(t *testing.T) {
	arrived := make(chan struct{})
	release := make(chan struct{})
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7",
				func(
					_ http.ResponseWriter,
					q *http.Request,
				) {
					close(arrived)

					select {
					case <-q.Context().Done():
					case <-release:
					}
				},
			)
		},
	)
	t.Cleanup(func() { close(release) })
	x, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() {
		_, e := r.Client.Project(x, 7)
		done <- e
	}()
	<-arrived
	cancel()

	select {
	case e := <-done:
		assert.True(t, errors.Is(e, context.Canceled))
	case <-time.After(5 * time.Second):
		t.Fatal("the call kept waiting after its context was cancelled")
	}
}
