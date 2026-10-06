package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestGetSendsBasicAuthToTheLocator(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"total":3}`)
	c := newJiraClient(t, s.URL)
	var out map[string]int
	assert.FatalOnError(
		t,
		c.Get(
			c.Base().Copy().Base(atlassian.JiraBase).Path("search/jql").Set(
				atlassian.JiraQueryKey,
				"project = ALFA",
			).String(),
			&out,
		),
	)
	assert.Integer(t, 3, out["total"])
	r := last()
	assert.String(t, "/rest/api/3/search/jql", r.URL.Path)
	assert.String(
		t,
		"project = ALFA",
		r.URL.Query().Get(atlassian.JiraQueryKey),
	)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-token", password)
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
}

func TestGetPathJoinsTheHost(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"issueLinkTypes":[]}`)
	var out map[string]any
	assert.FatalOnError(
		t,
		newJiraClient(t, s.URL).GetPath("rest/api/2/issueLinkType", &out),
	)
	assert.String(t, "/rest/api/2/issueLinkType", last().URL.Path)
}

func TestRefusalIsNotReadAsAnswer(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"errorMessages":["You are not authenticated."],"errors":{}}`,
	)
	c := newJiraClient(t, s.URL)
	var out map[string]any
	e := c.Get(c.Base().Copy().Path("rest/api/2/issue/ALFA-1").String(), &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"status: 401: You are not authenticated.",
		e.Error(),
	)
	assert.Count(t, 0, out)
}

func TestMissingIssueIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusNotFound,
		`{"errorMessages":["Issue does not exist"],"errors":{}}`,
	)
	c := newJiraClient(t, s.URL)
	_, e := c.Bytes(c.Base().Copy().Path("rest/api/3/issue/ALFA-99").String())
	assert.True(t, not_found.Is(e))
}

func TestBytesReturnsTheRawAnswer(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusOK, `{"key":"ALFA-1"}`)
	c := newJiraClient(t, s.URL)
	b, e := c.Bytes(c.Base().Copy().Path("rest/api/3/issue/ALFA-1").String())
	assert.FatalOnError(t, e)
	assert.String(t, `{"key":"ALFA-1"}`, string(b))
}
