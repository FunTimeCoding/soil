package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"testing"
)

func TestAMissingOwnerIsNotFound(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusNotFound,
		nil,
		`{"message":"Not Found"}`,
	).Repositories("alfa")
	assert.True(t, not_found.Is(e))
}

func TestAMissingRepositoryIsNotFound(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusNotFound,
		nil,
		`{"message":"Not Found"}`,
	).LatestRuns("alfa", "bravo")
	assert.True(t, not_found.Is(e))
}

func TestARateLimitIsUnavailable(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusForbidden,
		http.Header{
			"X-Ratelimit-Limit":     {"60"},
			"X-Ratelimit-Remaining": {"0"},
			"X-Ratelimit-Reset":     {"4102444800"},
		},
		`{"message":"API rate limit exceeded"}`,
	).LatestRuns("alfa", "bravo")
	assert.True(t, unavailable.Is(e))
}

func TestASecondaryRateLimitIsUnavailable(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusForbidden,
		http.Header{"Retry-After": {"60"}},
		`{"message":"You have exceeded a secondary rate limit","documentation_url":"https://docs.github.com/rest/overview/rate-limits-for-the-rest-api#about-secondary-rate-limits"}`,
	).SearchCode("alfa")
	assert.True(t, unavailable.Is(e))
}

func TestAMissingRepositoryHasNoTags(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusNotFound,
		nil,
		`{"message":"Not Found"}`,
	).Tags("alfa", "bravo")
	assert.True(t, not_found.Is(e))
}

func TestAConflictingRunDeleteIsConflict(t *testing.T) {
	assert.True(
		t,
		conflict.Is(
			refusingGithub(
				t,
				http.StatusConflict,
				nil,
				`{"message":"Cannot delete a workflow run that is in progress"}`,
			).DeleteRun("alfa", "bravo", 1),
		),
	)
}

func TestAnOutageFailsTheProbe(t *testing.T) {
	assert.True(
		t,
		unavailable.Is(
			refusingGithub(
				t,
				http.StatusServiceUnavailable,
				nil,
				`{"message":"Service Unavailable"}`,
			).Reachable(t.Context()),
		),
	)
}

func TestARefusedTokenIsOursToLookAt(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusUnauthorized,
		nil,
		`{"message":"Bad credentials"}`,
	).User()
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestAForbiddenDeleteIsOursToLookAt(t *testing.T) {
	e := refusingGithub(
		t,
		http.StatusForbidden,
		nil,
		`{"message":"Resource not accessible by personal access token"}`,
	).DeleteTag("alfa", "bravo", "v1.0.0")
	assert.True(t, classify.Reportable(e))
	assert.False(t, forbidden.Is(e))
}

func TestAnInvalidSearchIsOursToLookAt(t *testing.T) {
	_, e := refusingGithub(
		t,
		http.StatusUnprocessableEntity,
		nil,
		`{"message":"Validation Failed"}`,
	).SearchRepository("alfa")
	assert.True(t, classify.Reportable(e))
	assert.False(t, validation.Is(e))
}
