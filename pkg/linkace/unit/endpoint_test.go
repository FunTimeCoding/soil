package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"testing"
)

func TestPostSendsNotationWithToken(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusCreated, `{"id":7}`)
	var out map[string]int
	e := newClient(t, s.URL).Post(
		"links",
		map[string]string{"url": "https://bravo.example"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 7, out["id"])
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/api/v2/links", r.URL.Path)
	assert.String(t, "Bearer alfa-token", r.Header.Get(constant.Authorization))
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)
	assert.String(t, `{"url":"https://bravo.example"}`, string(b))
}

func TestPatchUsesItsVerb(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"id":7}`)
	var out map[string]int
	e := newClient(t, s.URL).Patch(
		"links/7",
		map[string]string{"title": "charlie"},
		&out,
	)
	assert.FatalOnError(t, e)
	assert.String(t, http.MethodPatch, last().Method)
}

func TestDeleteAcceptsNoContent(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusNoContent, "")
	assert.FatalOnError(t, newClient(t, s.URL).Delete("links/7"))
	assert.String(t, http.MethodDelete, last().Method)
}

func TestMissingLinkIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusNotFound, "")
	var out map[string]any
	e := newClient(t, s.URL).Get("links/99", nil, &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesLinkAceMessage(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnprocessableEntity,
		`{"message":"The url field is required."}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Post("links", map[string]string{}, &out)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "The url field is required.", e.Error())
}

func TestARefusedTokenIsForbidden(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"message":"Unauthenticated."}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Get("lists", nil, &out)
	assert.True(t, forbidden.Is(e))
}

func TestAForbiddenActionIsForbidden(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusForbidden,
		`{"message":"This action is unauthorized."}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Get("lists", nil, &out)
	assert.True(t, forbidden.Is(e))
}

func TestAConflictIsConflict(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusConflict,
		`{"message":"The tag already exists."}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Post("tags", map[string]string{}, &out)
	assert.True(t, conflict.Is(e))
}

func TestAMaintenanceAnswerIsUnavailable(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusServiceUnavailable,
		`{"message":"Service Unavailable"}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Post("lists", map[string]string{}, &out)
	assert.True(t, unavailable.Is(e))
}

func TestAnyOtherRefusalStaysUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusInternalServerError,
		`{"message":"Server Error"}`,
	)
	var out map[string]any
	e := newClient(t, s.URL).Get("lists", nil, &out)
	assert.True(t, unexpected.Is(e))
}
