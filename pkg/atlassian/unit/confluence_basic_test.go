package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/unit/atlassian_tester"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"testing"
)

func TestGetV2PathSendsBasicAuth(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"results":[]}`))
	var out map[string]any
	assert.FatalOnError(
		t,
		atlassian_tester.New(t).Confluence(s.URL).GetV2Path(
			constant.ConfluenceLabel,
			&out,
		),
	)
	r := s.Last()
	assert.String(t, "/wiki/api/v2/labels", r.URL.Path)
	user, password, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa", user)
	assert.String(t, "bravo-token", password)
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
}

func TestGetPathUsesOldBase(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{}`))
	var out map[string]any
	assert.FatalOnError(
		t,
		atlassian_tester.New(t).Confluence(s.URL).GetPath(
			constant.ConfluenceUser,
			&out,
		),
	)
	assert.String(t, "/wiki/rest/api/user/current", s.Last().URL.Path)
}

func TestNextFollowsCursorWithCredentials(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"results":[]}`))
	c := atlassian_tester.New(t).Confluence(s.URL)
	defer c.Close()
	var out map[string]any
	assert.FatalOnError(
		t,
		c.Get(c.Next("/wiki/api/v2/spaces?cursor=charlie"), &out),
	)
	r := s.Last()
	assert.String(t, "/wiki/api/v2/spaces", r.URL.Path)
	assert.String(t, "charlie", r.URL.Query().Get("cursor"))
	_, _, okay := r.BasicAuth()
	assert.True(t, okay)
}

func TestPostV2PathSendsBody(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"id":"7"}`))
	var out map[string]string
	assert.FatalOnError(
		t,
		atlassian_tester.New(t).Confluence(s.URL).PostV2Path(
			constant.ConfluencePage,
			`{"title":"delta"}`,
			&out,
		),
	)
	assert.String(t, "7", out["id"])
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "application/json", r.Header.Get(webConstant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"title":"delta"}`, string(b))
}

func TestPutV2PathUsesItsVerb(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"id":"7"}`))
	var out map[string]string
	assert.FatalOnError(
		t,
		atlassian_tester.New(t).Confluence(s.URL).PutV2Path(
			"/pages/7",
			`{"title":"echo"}`,
			&out,
		),
	)
	assert.String(t, http.MethodPut, s.Last().Method)
}

func TestPostOldPathAcceptsEmptyAnswer(t *testing.T) {
	s := upstream_server.New(t)
	assert.FatalOnError(
		t,
		atlassian_tester.New(t).Confluence(s.URL).PostOldPath(
			"/content",
			`{"type":"comment"}`,
		),
	)
	assert.String(t, "/wiki/rest/api/content", s.Last().URL.Path)
}

func TestDeleteAcceptsNoContent(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNoContent),
	)
	c := atlassian_tester.New(t).Confluence(s.URL)
	defer c.Close()
	assert.FatalOnError(t, c.Delete(c.Base().Copy().Path("/pages/7").String()))
	r := s.Last()
	assert.String(t, http.MethodDelete, r.Method)
	assert.String(t, "/wiki/api/v2/pages/7", r.URL.Path)
}

func TestMissingPageIsNotFound(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	var out map[string]any
	e := atlassian_tester.New(t).Confluence(s.URL).GetV2Path("/pages/99", &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesVersionTwoTitle(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"errors":[{"status":400,"title":"Invalid space identifier"}]}`,
		),
	)
	var out map[string]any
	e := atlassian_tester.New(t).Confluence(s.URL).GetV2Path(
		constant.ConfluenceSpace,
		&out,
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 400: Invalid space identifier", e.Error())
}

func TestRefusalCarriesOldMessage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"statusCode":403,"message":"Not permitted to use confluence"}`,
		),
	)
	var out map[string]any
	e := atlassian_tester.New(t).Confluence(s.URL).GetPath(
		constant.ConfluenceUser,
		&out,
	)
	assert.True(t, forbidden.Is(e))
	assert.StringContains(t, "Not permitted to use confluence", e.Error())
}
