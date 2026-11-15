package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"github.com/funtimecoding/soil/pkg/web/unit/requester_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
	"time"
)

func TestPaceSpacesRequests(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("alfa"))
	r := requester_tester.New(t, s.URL).WithPace(40 * time.Millisecond)
	start := time.Now()

	for range 3 {
		_, e := r.Bytes(request.Get("/items"))
		assert.FatalOnError(t, e)
	}

	assert.True(t, time.Since(start) >= 80*time.Millisecond)
	assert.Integer(t, 3, s.Count())
}

func TestPaceSpacesRetry(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(
			http.StatusServiceUnavailable,
			http.StatusOK,
		),
		upstream_server.WithAnswer("bravo"),
	)
	r := requester_tester.New(t, s.URL).WithPace(40 * time.Millisecond)
	start := time.Now()
	_, e := r.Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	assert.True(t, time.Since(start) >= 40*time.Millisecond)
	assert.Integer(t, 2, s.Count())
}

func TestPaceEndsWithItsContext(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("charlie"))
	r := requester_tester.New(t, s.URL).WithPace(time.Hour)
	_, e := r.Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	x, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, e = r.Bytes(request.Get("/items").WithContext(x))
	assert.True(t, errors.Deadline(e))
	assert.Integer(t, 1, s.Count())
}
