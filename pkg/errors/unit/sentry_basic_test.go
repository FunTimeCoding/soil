package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unit/sentry_tester"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsTokenToTrailingPath(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`[{"id":"7"}]`))
	var out []map[string]string
	e := sentry_tester.New(t, s.URL).Get(
		"organizations/alfa/issues",
		map[string]string{"query": "is:unresolved"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "7", out[0]["id"])
	r := s.Last()
	assert.String(t, http.MethodGet, r.Method)
	assert.String(t, "/api/0/organizations/alfa/issues/", r.URL.Path)
	assert.String(t, "is:unresolved", r.URL.Query().Get("query"))
	assert.String(
		t,
		"Bearer alfa-token",
		r.Header.Get(webConstant.Authorization),
	)
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
}

func TestGetValuesKeepsRepeatedFields(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"data":[]}`))
	var out map[string]any
	e := sentry_tester.New(t, s.URL).GetValues(
		"organizations/alfa/events",
		url.Values{"field": {"id", "title"}},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"id", "title"}, s.Last().URL.Query()["field"])
}

func TestPutSendsNotation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"status":"resolved"}`),
	)
	var out map[string]string
	e := sentry_tester.New(t, s.URL).Put(
		"organizations/alfa/issues/7",
		map[string]string{"status": "resolved"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "resolved", out[constant.Status])
	r := s.Last()
	assert.String(t, http.MethodPut, r.Method)
	assert.String(t, "/api/0/organizations/alfa/issues/7/", r.URL.Path)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)
	assert.String(t, `{"status":"resolved"}`, string(b))
}

func TestDeleteAcceptsAccepted(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusAccepted))
	assert.FatalOnError(
		t,
		sentry_tester.New(t, s.URL).Delete("organizations/alfa/issues/7"),
	)
	assert.String(t, http.MethodDelete, s.Last().Method)
}

func TestMissingIssueIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"detail":"The requested resource does not exist"}`,
		),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get(
		"organizations/alfa/issues/99",
		nil,
		&out,
	)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesSentryDetail(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"detail":"You do not have permission to perform this action."}`,
		),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"You do not have permission to perform this action.",
		e.Error(),
	)
}

func TestRejectedQueryIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(`{"detail":"Invalid search query"}`),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, validation.Is(e))
}

func TestExpiredTokenIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"detail":"Invalid token"}`),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get("auth", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestConflictingUpdateIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(`{"detail":"Issue is locked"}`),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Put(
		"organizations/alfa/issues/7",
		map[string]string{"status": "resolved"},
		&out,
	)
	assert.True(t, conflict.Is(e))
}

func TestAnyOtherRefusalStaysUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(`{"detail":"Internal error"}`),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestAnswerThatIsNotNotationIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer("<html>login</html>"),
	)
	var out map[string]any
	e := sentry_tester.New(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, unexpected.Is(e))
}
