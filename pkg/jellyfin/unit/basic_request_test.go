package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/jellyfin/unit/jellyfin_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsTokenAndParameters(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"Name":"bravo"}`))
	var out map[string]string
	assert.FatalOnError(
		t,
		jellyfin_tester.Basic(t, s.URL).Get(
			"/Items/charlie",
			url.Values{"Fields": {"Name"}},
			&out,
		),
	)
	assert.String(t, "bravo", out["Name"])
	r := s.Last()
	assert.String(t, "/Items/charlie", r.URL.Path)
	assert.String(t, "Name", r.URL.Query().Get("Fields"))
	assert.String(
		t,
		`MediaBrowser Token="alfa-token"`,
		r.Header.Get(webConstant.Authorization),
	)
}

func TestPostSendsParametersAndNotation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNoContent),
	)
	assert.FatalOnError(
		t,
		jellyfin_tester.Basic(t, s.URL).Post(
			"/Sessions/delta/Command",
			url.Values{"seekPositionTicks": {"5"}},
			map[string]string{"Name": "SetVolume"},
		),
	)
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "5", r.URL.Query().Get("seekPositionTicks"))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"Name":"SetVolume"}`, string(b))
}

func TestPostWithoutBodySendsNothing(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNoContent),
	)
	assert.FatalOnError(
		t,
		jellyfin_tester.Basic(t, s.URL).Post(
			"/Sessions/echo/Playing",
			nil,
			nil,
		),
	)
	assert.String(t, "", s.LastBody())
}

func TestBadRequestIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"title":"One or more validation errors occurred.","status":400}`,
		),
	)
	var out map[string]string
	e := jellyfin_tester.Basic(t, s.URL).Rejecting().Get(
		"/Items/foxtrot",
		nil,
		&out,
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "One or more validation errors occurred.", e.Error())
}

func TestBadRequestElsewhereIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"title":"One or more validation errors occurred.","status":400}`,
		),
	)
	var out map[string]string
	e := jellyfin_tester.Basic(t, s.URL).Get("/Items", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestUnauthorizedRequestIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"title":"Unauthorized","status":401}`),
	)
	var out map[string]string
	e := jellyfin_tester.Basic(t, s.URL).Get("/Items/foxtrot", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestMissingItemIsNotFound(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	var out map[string]string
	e := jellyfin_tester.Basic(t, s.URL).Get("/Items/golf", nil, &out)
	assert.True(t, not_found.Is(e))
}
