package unit

import (
	"net/http"
	"net/http/httptest"
)

func served(
	m *http.ServeMux,
	path string,
) *httptest.ResponseRecorder {
	result := httptest.NewRecorder()
	m.ServeHTTP(result, httptest.NewRequest(http.MethodGet, path, nil))

	return result
}
