package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"github.com/funtimecoding/soil/pkg/web/unit/requester_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
	"time"
)

func TestSinceSendsModifiedHeaderAndAcceptsNotModified(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotModified),
	)
	since := time.Date(2026, time.November, 19, 8, 52, 16, 0, time.UTC)
	r, e := requester_tester.New(t, s.URL).Send(
		request.Get("/document").WithSince(since),
	)
	assert.FatalOnError(t, e)
	defer errors.LogClose(r.Body)
	assert.Integer(t, http.StatusNotModified, r.StatusCode)
	assert.String(
		t,
		"Thu, 19 Nov 2026 08:52:16 GMT",
		s.Last().Header.Get(constant.IfModifiedSince),
	)
}

func TestNotModifiedWithoutSinceStaysRefusal(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotModified),
	)
	_, e := requester_tester.New(t, s.URL).Send(request.Get("/document"))
	assert.True(t, unexpected.Is(e))
	assert.String(t, "", s.Last().Header.Get(constant.IfModifiedSince))
}
