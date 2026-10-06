package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"testing"
)

func TestGetV2PathSendsBasicAuth(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"results":[]}`)
	var out map[string]any
	assert.FatalOnError(
		t,
		newConfluenceClient(t, s.URL).GetV2Path(atlassian.ConfluenceLabel, &out),
	)
	r := last()
	assert.String(t, "/wiki/api/v2/labels", r.URL.Path)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-token", password)
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
}

func TestGetPathUsesTheOldBase(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{}`)
	var out map[string]any
	assert.FatalOnError(
		t,
		newConfluenceClient(t, s.URL).GetPath(atlassian.ConfluenceUser, &out),
	)
	assert.String(t, "/wiki/rest/api/user/current", last().URL.Path)
}

func TestNextFollowsTheCursorWithCredentials(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"results":[]}`)
	c := newConfluenceClient(t, s.URL)
	var out map[string]any
	assert.FatalOnError(
		t,
		c.Get(c.Next("/wiki/api/v2/spaces?cursor=charlie"), &out),
	)
	r := last()
	assert.String(t, "/wiki/api/v2/spaces", r.URL.Path)
	assert.String(t, "charlie", r.URL.Query().Get("cursor"))
	_, _, okay := r.BasicAuth()
	assert.True(t, okay)
}

func TestPostV2PathSendsTheBody(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"id":"7"}`)
	var out map[string]string
	assert.FatalOnError(
		t,
		newConfluenceClient(t, s.URL).PostV2Path(
			atlassian.ConfluencePage,
			`{"title":"delta"}`,
			&out,
		),
	)
	assert.String(t, "7", out["id"])
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "application/json", r.Header.Get(constant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"title":"delta"}`, string(b))
}

func TestPutV2PathUsesItsVerb(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"id":"7"}`)
	var out map[string]string
	assert.FatalOnError(
		t,
		newConfluenceClient(t, s.URL).PutV2Path(
			"/pages/7",
			`{"title":"echo"}`,
			&out,
		),
	)
	assert.String(t, http.MethodPut, last().Method)
}

func TestPostOldPathAcceptsAnEmptyAnswer(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, "")
	assert.FatalOnError(
		t,
		newConfluenceClient(t, s.URL).PostOldPath(
			"/content",
			`{"type":"comment"}`,
		),
	)
	assert.String(t, "/wiki/rest/api/content", last().URL.Path)
}

func TestDeleteAcceptsNoContent(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusNoContent, "")
	c := newConfluenceClient(t, s.URL)
	assert.FatalOnError(t, c.Delete(c.Base().Copy().Path("/pages/7").String()))
	r := last()
	assert.String(t, http.MethodDelete, r.Method)
	assert.String(t, "/wiki/api/v2/pages/7", r.URL.Path)
}

func TestMissingPageIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusNotFound, "")
	var out map[string]any
	e := newConfluenceClient(t, s.URL).GetV2Path("/pages/99", &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesTheVersionTwoTitle(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusBadRequest,
		`{"errors":[{"status":400,"title":"Invalid space identifier"}]}`,
	)
	var out map[string]any
	e := newConfluenceClient(t, s.URL).GetV2Path(
		atlassian.ConfluenceSpace,
		&out,
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 400: Invalid space identifier", e.Error())
}

func TestRefusalCarriesTheOldMessage(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusForbidden,
		`{"statusCode":403,"message":"Not permitted to use confluence"}`,
	)
	var out map[string]any
	e := newConfluenceClient(t, s.URL).GetPath(atlassian.ConfluenceUser, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"status: 403: Not permitted to use confluence",
		e.Error(),
	)
}
