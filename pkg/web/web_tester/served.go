package web_tester

import (
	"net/http"
	"net/http/httptest"
)

func Served(
	m *http.ServeMux,
	path string,
) *httptest.ResponseRecorder {
	result := httptest.NewRecorder()
	m.ServeHTTP(result, httptest.NewRequest(http.MethodGet, path, nil))

	return result
}
