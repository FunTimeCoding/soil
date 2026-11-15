package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestMissingProjectIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"message":"404 Project Not Found"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, not_found.Is(e))
	assert.String(t, "Not Found", e.Error())
}

func TestForbiddenProjectIsForbidden(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(`{"message":"403 Forbidden"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, forbidden.Is(e))
}

func TestUnprocessableRequestIsCallersMistake(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnprocessableEntity),
		upstream_server.WithAnswer(`{"message":"name is taken"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, validation.Is(e))
}

func TestFailedPreconditionIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusPreconditionFailed),
		upstream_server.WithAnswer(`{"message":"modified since"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, conflict.Is(e))
}

func TestRateLimitedRequestIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusTooManyRequests),
		upstream_server.WithAnswer(`{"message":"slow down"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, unavailable.Is(e))
}

func TestOutagePageIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadGateway),
		upstream_server.WithAnswer("GitLab is not responding"),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, unavailable.Is(e))
}

func TestUnauthorizedTokenIsOursToLookAt(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"message":"401 Unauthorized"}`),
	)
	_, e := client_tester.Connect(t, s.URL).Project(t.Context(), 7)
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestRefusedDialIsConnectionFailure(t *testing.T) {
	_, e := client_tester.Connect(t, outage_tester.RefusedAddress(t)).Project(
		t.Context(),
		7,
	)
	var f *connection.Failure
	assert.True(t, errors.As(e, &f))
}
