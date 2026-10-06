package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsTokenToTrailingPath(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `[{"id":"7"}]`)
	var out []map[string]string
	e := newSentryClient(t, s.URL).Get(
		"organizations/alfa/issues",
		map[string]string{"query": "is:unresolved"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "7", out[0]["id"])
	r := last()
	assert.String(t, http.MethodGet, r.Method)
	assert.String(t, "/api/0/organizations/alfa/issues/", r.URL.Path)
	assert.String(t, "is:unresolved", r.URL.Query().Get("query"))
	assert.String(t, "Bearer alfa-token", r.Header.Get(constant.Authorization))
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
}

func TestGetValuesKeepsRepeatedFields(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"data":[]}`)
	var out map[string]any
	e := newSentryClient(t, s.URL).GetValues(
		"organizations/alfa/events",
		url.Values{"field": {"id", "title"}},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"id", "title"}, last().URL.Query()["field"])
}

func TestPutSendsNotation(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"status":"resolved"}`)
	var out map[string]string
	e := newSentryClient(t, s.URL).Put(
		"organizations/alfa/issues/7",
		map[string]string{"status": "resolved"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "resolved", out[errorsConstant.Status])
	r := last()
	assert.String(t, http.MethodPut, r.Method)
	assert.String(t, "/api/0/organizations/alfa/issues/7/", r.URL.Path)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)
	assert.String(t, `{"status":"resolved"}`, string(b))
}

func TestDeleteAcceptsAccepted(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusAccepted, "")
	assert.FatalOnError(
		t,
		newSentryClient(t, s.URL).Delete("organizations/alfa/issues/7"),
	)
	assert.String(t, http.MethodDelete, last().Method)
}

func TestMissingIssueIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusNotFound,
		`{"detail":"The requested resource does not exist"}`,
	)
	var out map[string]any
	e := newSentryClient(t, s.URL).Get(
		"organizations/alfa/issues/99",
		nil,
		&out,
	)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesSentryDetail(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusForbidden,
		`{"detail":"You do not have permission to perform this action."}`,
	)
	var out map[string]any
	e := newSentryClient(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"status: 403: You do not have permission to perform this action.",
		e.Error(),
	)
}

func TestAnswerThatIsNotNotationIsUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusOK, "<html>login</html>")
	var out map[string]any
	e := newSentryClient(t, s.URL).Get("organizations/alfa/issues", nil, &out)
	assert.True(t, unexpected.Is(e))
}
