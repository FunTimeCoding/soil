package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
	"testing"
	"time"
)

func TestPaceSpacesRequests(t *testing.T) {
	s, seen := newScriptedServer(t, "alfa", http.StatusOK)
	r := newRequester(t, s.URL).WithPace(40 * time.Millisecond)
	start := time.Now()

	for range 3 {
		_, e := r.Bytes(request.Get("/items"))
		assert.FatalOnError(t, e)
	}

	assert.True(t, time.Since(start) >= 80*time.Millisecond)
	assert.Count(t, 3, seen())
}

func TestPaceSpacesARetry(t *testing.T) {
	s, seen := newScriptedServer(
		t,
		"bravo",
		http.StatusServiceUnavailable,
		http.StatusOK,
	)
	r := newRequester(t, s.URL).WithPace(40 * time.Millisecond)
	start := time.Now()
	_, e := r.Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	assert.True(t, time.Since(start) >= 40*time.Millisecond)
	assert.Count(t, 2, seen())
}

func TestPaceEndsWithItsContext(t *testing.T) {
	s, seen := newScriptedServer(t, "charlie", http.StatusOK)
	r := newRequester(t, s.URL).WithPace(time.Hour)
	_, e := r.Bytes(request.Get("/items"))
	assert.FatalOnError(t, e)
	x, cancel := context.WithTimeout(
		context.Background(),
		50*time.Millisecond,
	)
	defer cancel()
	_, e = r.Bytes(request.Get("/items").WithContext(x))
	assert.True(t, errors.Deadline(e))
	assert.Count(t, 1, seen())
}
