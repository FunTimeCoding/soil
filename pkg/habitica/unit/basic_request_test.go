package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsCredentialsAndUnwrapsTheEnvelope(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"success":true,"data":[{"name":"charlie"}]}`,
	)
	var out []map[string]string
	assert.FatalOnError(
		t,
		newClient(t, s.URL).Get(
			"/tasks/user",
			url.Values{constant.TypeParameter: {"todos"}},
			&out,
		),
	)
	assert.String(t, "charlie", out[0]["name"])
	r := last()
	assert.String(t, "/api/v3/tasks/user", r.URL.Path)
	assert.String(t, "todos", r.URL.Query().Get(constant.TypeParameter))
	assert.String(t, "alfa-user", r.Header.Get(constant.UserHeader))
	assert.String(t, "bravo-token", r.Header.Get(constant.TokenHeader))
}

func TestPostSendsParametersAndNotation(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"success":true,"data":{"str":3}}`,
	)
	var out map[string]int
	assert.FatalOnError(
		t,
		newClient(t, s.URL).Post(
			"/user/allocate",
			url.Values{constant.StatParameter: {"str"}},
			map[string]string{"text": "delta"},
			&out,
		),
	)
	assert.Integer(t, 3, out["str"])
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/api/v3/user/allocate", r.URL.Path)
	assert.String(t, "str", r.URL.Query().Get(constant.StatParameter))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"text":"delta"}`, string(b))
}

func TestPostDiscardIgnoresTheAnswer(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, "not notation")
	assert.FatalOnError(t, newClient(t, s.URL).PostDiscard("/cron"))
	assert.String(t, "/api/v3/cron", last().URL.Path)
}

func TestRefusalCarriesTheMessage(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"success":false,"error":"NotAuthorized","message":"echo refused"}`,
	)
	var out map[string]string
	e := newClient(t, s.URL).Get("/user", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "echo refused", e.Error())
}

func TestMissingTaskIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusNotFound,
		`{"success":false,"error":"NotFound","message":"Task not found."}`,
	)
	var out map[string]string
	e := newClient(t, s.URL).Post("/tasks/foxtrot/score/up", nil, nil, &out)
	assert.True(t, not_found.Is(e))
}
