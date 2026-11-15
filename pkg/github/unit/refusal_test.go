package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/github/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestMissingOwnerIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"message":"Not Found"}`),
	)
	_, e := client_tester.New(t, s.URL).Repositories("alfa")
	assert.True(t, not_found.Is(e))
}

func TestMissingRepositoryIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"message":"Not Found"}`),
	)
	_, e := client_tester.New(t, s.URL).LatestRuns("alfa", "bravo")
	assert.True(t, not_found.Is(e))
}

func TestRateLimitIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithHeader(
			http.Header{
				"X-Ratelimit-Limit":     {"60"},
				"X-Ratelimit-Remaining": {"0"},
				"X-Ratelimit-Reset":     {"4102444800"},
			},
		),
		upstream_server.WithAnswer(`{"message":"API rate limit exceeded"}`),
	)
	_, e := client_tester.New(t, s.URL).LatestRuns("alfa", "bravo")
	assert.True(t, unavailable.Is(e))
}

func TestSecondaryRateLimitIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithHeader(http.Header{"Retry-After": {"60"}}),
		upstream_server.WithAnswer(
			`{"message":"You have exceeded a secondary rate limit","documentation_url":"https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits"}`,
		),
	)
	_, e := client_tester.New(t, s.URL).SearchCode("alfa")
	assert.True(t, unavailable.Is(e))
}

func TestMissingRepositoryHasNoTags(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"message":"Not Found"}`),
	)
	_, e := client_tester.New(t, s.URL).Tags("alfa", "bravo")
	assert.True(t, not_found.Is(e))
}

func TestConflictingRunDeleteIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(
			`{"message":"Cannot delete a workflow run that is in progress"}`,
		),
	)
	assert.True(
		t,
		conflict.Is(client_tester.New(t, s.URL).DeleteRun("alfa", "bravo", 1)),
	)
}

func TestOutageFailsProbe(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"message":"Service Unavailable"}`),
	)
	assert.True(
		t,
		unavailable.Is(client_tester.New(t, s.URL).Reachable(t.Context())),
	)
}

func TestRefusedTokenIsOursToLookAt(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"message":"Bad credentials"}`),
	)
	_, e := client_tester.New(t, s.URL).User()
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestForbiddenDeleteIsOursToLookAt(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"message":"Resource not accessible by personal access token"}`,
		),
	)
	e := client_tester.New(t, s.URL).DeleteTag("alfa", "bravo", "v1.0.0")
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestInvalidSearchIsOursToLookAt(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnprocessableEntity),
		upstream_server.WithAnswer(`{"message":"Validation Failed"}`),
	)
	_, e := client_tester.New(t, s.URL).SearchRepository("alfa")
	assert.True(t, classify.Reportable(e))
	assert.False(t, validation.Is(e))
}
