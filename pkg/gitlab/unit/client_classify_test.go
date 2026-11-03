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
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func TestAMissingProjectIsNotFound(t *testing.T) {
	_, e := refusingProject(t, http.StatusNotFound, "404 Project Not Found").
		Client.Project(t.Context(), 7)
	assert.True(t, not_found.Is(e))
	assert.String(t, "Not Found", e.Error())
}

func TestAForbiddenProjectIsForbidden(t *testing.T) {
	_, e := refusingProject(t, http.StatusForbidden, "403 Forbidden").
		Client.Project(t.Context(), 7)
	assert.True(t, forbidden.Is(e))
}

func TestAnUnprocessableRequestIsTheCallersMistake(t *testing.T) {
	_, e := refusingProject(t, http.StatusUnprocessableEntity, "name is taken").
		Client.Project(t.Context(), 7)
	assert.True(t, validation.Is(e))
}

func TestAFailedPreconditionIsAConflict(t *testing.T) {
	_, e := refusingProject(t, http.StatusPreconditionFailed, "modified since").
		Client.Project(t.Context(), 7)
	assert.True(t, conflict.Is(e))
}

func TestARateLimitedRequestIsUnavailable(t *testing.T) {
	_, e := refusingProject(t, http.StatusTooManyRequests, "slow down").
		Client.Project(t.Context(), 7)
	assert.True(t, unavailable.Is(e))
}

func TestAnOutagePageIsUnavailable(t *testing.T) {
	_, e := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7",
				func(
					w http.ResponseWriter,
					_ *http.Request,
				) {
					web.TextHeader(w)
					web.Write(
						w,
						http.StatusBadGateway,
						"GitLab is not responding",
					)
				},
			)
		},
	).Client.Project(t.Context(), 7)
	assert.True(t, unavailable.Is(e))
}

func TestAnUnauthorizedTokenIsOursToLookAt(t *testing.T) {
	_, e := refusingProject(t, http.StatusUnauthorized, "401 Unauthorized").
		Client.Project(t.Context(), 7)
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestARefusedDialIsAConnectionFailure(t *testing.T) {
	c := gitlab.New(
		"127.0.0.1:1",
		"test-token",
		gitlab.WithInsecure(),
		gitlab.WithRetryMax(0),
	)
	t.Cleanup(c.Close)
	_, e := c.Project(t.Context(), 7)
	var f *connection.Failure
	assert.True(t, errors.As(e, &f))
}
