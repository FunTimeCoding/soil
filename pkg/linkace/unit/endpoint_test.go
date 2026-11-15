package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/linkace/unit/linkace_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"testing"
)

func TestPostSendsNotationWithToken(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusCreated),
		upstream_server.WithAnswer(`{"id":7}`),
	)
	var out map[string]int
	e := linkace_tester.New(t, s.URL).Post(
		"links",
		map[string]string{"url": "https://bravo.example"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, out["id"])
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/api/v2/links", r.URL.Path)
	assert.String(
		t,
		"Bearer alfa-token",
		r.Header.Get(webConstant.Authorization),
	)
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)
	assert.String(t, `{"url":"https://bravo.example"}`, string(b))
}

func TestPatchUsesItsVerb(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"id":7}`))
	var out map[string]int
	e := linkace_tester.New(t, s.URL).Patch(
		"links/7",
		map[string]string{"title": "charlie"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, http.MethodPatch, s.Last().Method)
}

func TestDeleteAcceptsNoContent(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNoContent),
	)
	assert.FatalOnError(t, linkace_tester.New(t, s.URL).Delete("links/7"))
	assert.String(t, http.MethodDelete, s.Last().Method)
}

func TestMissingLinkIsNotFound(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Get("links/99", nil, &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesLinkAceMessage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnprocessableEntity),
		upstream_server.WithAnswer(`{"message":"The url field is required."}`),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Post("links", map[string]string{}, &out)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "The url field is required.", e.Error())
}

func TestRefusedTokenIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"message":"Unauthenticated."}`),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Get("lists", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestRefusedActionIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"message":"This action is unauthorized."}`,
		),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Get("lists", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestConflictIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(`{"message":"The tag already exists."}`),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Post("tags", map[string]string{}, &out)
	assert.True(t, conflict.Is(e))
}

func TestMaintenanceAnswerIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"message":"Service Unavailable"}`),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Post("lists", map[string]string{}, &out)
	assert.True(t, unavailable.Is(e))
}

func TestAnyOtherRefusalStaysUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(`{"message":"Server Error"}`),
	)
	var out map[string]any
	e := linkace_tester.New(t, s.URL).Get("lists", nil, &out)
	assert.True(t, unexpected.Is(e))
}
