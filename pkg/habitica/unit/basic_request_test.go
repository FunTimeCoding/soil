package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/habitica/constant"
	"github.com/funtimecoding/soil/pkg/habitica/unit/habitica_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetSendsCredentialsAndUnwrapsEnvelope(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"success":true,"data":[{"name":"charlie"}]}`,
		),
	)
	var out []map[string]string
	assert.FatalOnError(
		t,
		habitica_tester.New(t, s.URL).Get(
			"/tasks/user",
			url.Values{constant.TypeParameter: {"todos"}},
			&out,
		),
	)
	assert.String(t, "charlie", out[0]["name"])
	r := s.Last()
	assert.String(t, "/api/v3/tasks/user", r.URL.Path)
	assert.String(t, "todos", r.URL.Query().Get(constant.TypeParameter))
	assert.String(t, "alfa-user", r.Header.Get(constant.UserHeader))
	assert.String(t, "bravo-token", r.Header.Get(constant.TokenHeader))
}

func TestPostSendsParametersAndNotation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"success":true,"data":{"str":3}}`),
	)
	var out map[string]int
	assert.FatalOnError(
		t,
		habitica_tester.New(t, s.URL).Post(
			"/user/allocate",
			url.Values{constant.StatParameter: {"str"}},
			map[string]string{"text": "delta"},
			&out,
		),
	)
	assert.Integer(t, 3, out["str"])
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/api/v3/user/allocate", r.URL.Path)
	assert.String(t, "str", r.URL.Query().Get(constant.StatParameter))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"text":"delta"}`, string(b))
}

func TestPostDiscardIgnoresAnswer(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer("not notation"))
	assert.FatalOnError(t, habitica_tester.New(t, s.URL).PostDiscard("/cron"))
	assert.String(t, "/api/v3/cron", s.Last().URL.Path)
}

func TestRefusalCarriesMessage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(
			`{"success":false,"error":"NotAuthorized","message":"echo refused"}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Get("/user", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "echo refused", e.Error())
}

func TestForbiddenAnswerIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"success":false,"error":"Forbidden","message":"not yours"}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Get("/tasks/user", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestConflictIsConflict(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusConflict),
		upstream_server.WithAnswer(
			`{"success":false,"error":"Conflict","message":"cron already run"}`,
		),
	)
	e := habitica_tester.New(t, s.URL).PostDiscard("/cron")
	assert.True(t, conflict.Is(e))
}

func TestAnswerOtherThanClassesStaysUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
		upstream_server.WithAnswer(
			`{"success":false,"error":"InternalServerError","message":"boom"}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Get("/user", nil, &out)
	assert.True(t, unexpected.Is(e))
}

func TestMissingTaskIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(
			`{"success":false,"error":"NotFound","message":"Task not found."}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Post(
		"/tasks/foxtrot/score/up",
		nil,
		nil,
		&out,
	)
	assert.True(t, not_found.Is(e))
}

func TestBadRequestIsValidation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"success":false,"error":"BadRequest","message":"\"bravo\" is not a valid Stat."}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Rejecting().Post(
		"/user/allocate",
		nil,
		nil,
		&out,
	)
	assert.True(t, validation.Is(e))
	assert.String(t, `"bravo" is not a valid Stat.`, e.Error())
}

func TestBadRequestElsewhereIsFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"success":false,"error":"BadRequest","message":"malformed"}`,
		),
	)
	var out map[string]string
	e := habitica_tester.New(t, s.URL).Post("/tasks/user", nil, nil, &out)
	assert.True(t, unexpected.Is(e))
}
