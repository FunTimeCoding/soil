package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func tokenProtected(tokens []string) http.Handler {
	return web.TokenMiddleware(tokens, constant.HealthPath)(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(http.StatusOK)
			},
		),
	)
}

func tokenRequest(path string, token string) *httptest.ResponseRecorder {
	q := httptest.NewRequest(http.MethodGet, path, nil)

	if token != "" {
		web.Bearer(q, token)
	}

	w := httptest.NewRecorder()
	tokenProtected([]string{strings.UpperAlfa, strings.UpperBravo}).ServeHTTP(
		w,
		q,
	)

	return w
}

func TestTokenMiddlewareAccepts(t *testing.T) {
	assert.Integer(
		t,
		http.StatusOK,
		tokenRequest(constant.RootPath, strings.UpperAlfa).Code,
	)
	assert.Integer(
		t,
		http.StatusOK,
		tokenRequest(constant.RootPath, strings.UpperBravo).Code,
	)
}

func TestTokenMiddlewareRejects(t *testing.T) {
	assert.Integer(
		t,
		http.StatusUnauthorized,
		tokenRequest(constant.RootPath, "").Code,
	)
	assert.Integer(
		t,
		http.StatusUnauthorized,
		tokenRequest(constant.RootPath, strings.UpperCharlie).Code,
	)
}

func TestTokenMiddlewareExempt(t *testing.T) {
	assert.Integer(t, http.StatusOK, tokenRequest(constant.HealthPath, "").Code)
}

func TestTokenMiddlewareEmptySetRejects(t *testing.T) {
	q := httptest.NewRequest(http.MethodGet, constant.RootPath, nil)
	web.Bearer(q, strings.UpperAlfa)
	w := httptest.NewRecorder()
	tokenProtected(nil).ServeHTTP(w, q)
	assert.Integer(t, http.StatusUnauthorized, w.Code)
}

func TestTokenMiddlewareTokenSchemeRejected(t *testing.T) {
	q := httptest.NewRequest(http.MethodGet, constant.RootPath, nil)
	web.Token(q, strings.UpperAlfa)
	w := httptest.NewRecorder()
	tokenProtected([]string{strings.UpperAlfa}).ServeHTTP(w, q)
	assert.Integer(t, http.StatusUnauthorized, w.Code)
}
