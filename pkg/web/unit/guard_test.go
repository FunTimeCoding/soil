package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"net/http"
	"testing"
)

func TestTokenRejectsMissingBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", web_tester.Serve)
	assert.Integer(t, http.StatusUnauthorized, web_tester.Request(t, m, ""))
}

func TestTokenRejectsWrongBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", web_tester.Serve)
	assert.Integer(
		t,
		http.StatusUnauthorized,
		web_tester.Request(t, m, "bravo"),
	)
}

func TestTokenAcceptsBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Token("GET /target", web_tester.Serve)
	assert.Integer(t, http.StatusOK, web_tester.Request(t, m, "alfa"))
}

func TestTokenAcceptsRotationSibling(t *testing.T) {
	m := http.NewServeMux()
	g := guard.New(m, []string{"alfa", "bravo"})
	g.Token("GET /target", web_tester.Serve)
	assert.Integer(t, http.StatusOK, web_tester.Request(t, m, "alfa"))
	assert.Integer(t, http.StatusOK, web_tester.Request(t, m, "bravo"))
}

func TestOpenPassesWithoutBearer(t *testing.T) {
	m := http.NewServeMux()
	guard.New(m, []string{"alfa"}).Open("GET /target", web_tester.Serve)
	assert.Integer(t, http.StatusOK, web_tester.Request(t, m, ""))
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
	).Session("GET /target", web_tester.Serve)
	assert.Integer(t, http.StatusOK, web_tester.Request(t, m, ""))
	assert.True(t, called)
}

func TestSessionWithoutMiddlewarePanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	guard.New(http.NewServeMux(), []string{"alfa"}).Session(
		"GET /target",
		web_tester.Serve,
	)
}

func TestNewWithoutTokensPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	guard.New(http.NewServeMux(), nil)
}
