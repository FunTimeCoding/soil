package web_tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Request(
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
