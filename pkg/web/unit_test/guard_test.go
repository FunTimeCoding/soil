package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"net/http/httptest"
	"testing"
)

func serve(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusOK)
}

func request(
	t *testing.T,
	m *http.ServeMux,
	bearer string,
) int {
	t.Helper()
	q := httptest.NewRequest(http.MethodGet, "/target", nil)

	if bearer != "" {
		q.Header.Set(
			constant.Authorization,
			key_value.Space(constant.Bearer, bearer),
		)
	}

	w := httptest.NewRecorder()
	m.ServeHTTP(w, q)

	return w.Code
}

func TestTokenRejectsMissingBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", serve)
	assert.Integer(t, http.StatusUnauthorized, request(t, m, ""))
}

func TestTokenRejectsWrongBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", serve)
	assert.Integer(t, http.StatusUnauthorized, request(t, m, "bravo"))
}

func TestTokenAcceptsBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", serve)
	assert.Integer(t, http.StatusOK, request(t, m, "alfa"))
}

func TestTokenAcceptsRotationSibling(t *testing.T) {
	m := http.NewServeMux()
	g := guard.New(m, []string{"alfa", "bravo"})
	g.Token("GET /target", serve)
	assert.Integer(t, http.StatusOK, request(t, m, "alfa"))
	assert.Integer(t, http.StatusOK, request(t, m, "bravo"))
}

func TestOpenPassesWithoutBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Open("GET /target", serve)
	assert.Integer(t, http.StatusOK, request(t, m, ""))
}

func TestSessionUsesMiddleware(t *testing.T) {
	m := http.NewServeMux()
	called := false
	guard.New(m, []string{"alfa"}).WithSession(
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				called = true
				next(w, q)
			}
		},
	).Session("GET /target", serve)
	assert.Integer(t, http.StatusOK, request(t, m, ""))
	assert.True(t, called)
}

func TestSessionWithoutMiddlewarePanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	guard.New(http.NewServeMux(), []string{"alfa"}).Session(
		"GET /target",
		serve,
	)
}

func TestNewWithoutTokensPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	guard.New(http.NewServeMux(), nil)
}
