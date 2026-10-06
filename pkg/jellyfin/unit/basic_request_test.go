package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsTokenAndParameters(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"Name":"bravo"}`)
	var out map[string]string
	assert.FatalOnError(
		t,
		newClient(t, s.URL).Get(
			"/Items/charlie",
			url.Values{"Fields": {"Name"}},
			&out,
		),
	)
	assert.String(t, "bravo", out["Name"])
	r := last()
	assert.String(t, "/Items/charlie", r.URL.Path)
	assert.String(t, "Name", r.URL.Query().Get("Fields"))
	assert.String(
		t,
		`MediaBrowser Token="alfa-token"`,
		r.Header.Get(constant.Authorization),
	)
}

func TestPostSendsParametersAndNotation(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusNoContent, "")
	assert.FatalOnError(
		t,
		newClient(t, s.URL).Post(
			"/Sessions/delta/Command",
			url.Values{"seekPositionTicks": {"5"}},
			map[string]string{"Name": "SetVolume"},
		),
	)
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "5", r.URL.Query().Get("seekPositionTicks"))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"Name":"SetVolume"}`, string(b))
}

func TestPostWithoutBodySendsNothing(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusNoContent, "")
	assert.FatalOnError(
		t,
		newClient(t, s.URL).Post("/Sessions/echo/Playing", nil, nil),
	)
	b, e := io.ReadAll(last().Body)
	assert.FatalOnError(t, e)
	assert.String(t, "", string(b))
}

func TestRefusalCarriesProblemTitle(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusBadRequest,
		`{"title":"One or more validation errors occurred.","status":400}`,
	)
	var out map[string]string
	e := newClient(t, s.URL).Get("/Items/foxtrot", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"One or more validation errors occurred.",
		e.Error(),
	)
}

func TestMissingItemIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusNotFound, "")
	var out map[string]string
	e := newClient(t, s.URL).Get("/Items/golf", nil, &out)
	assert.True(t, not_found.Is(e))
}
