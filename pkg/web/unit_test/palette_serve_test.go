package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeEmptyQuery(t *testing.T) {
	serve := palette.NewServe(testRegistry())
	w := httptest.NewRecorder()
	q := httptest.NewRequest(http.MethodGet, "/palette", nil)
	serve(w, q)
	assert.Integer(t, http.StatusOK, w.Code)
	assert.StringContains(t, "Dashboard", w.Body.String())
	assert.StringContains(t, "Create project", w.Body.String())
	assert.StringContains(t, "palette-item", w.Body.String())
}

func TestServeWithQuery(t *testing.T) {
	serve := palette.NewServe(testRegistry())
	w := httptest.NewRecorder()
	q := httptest.NewRequest(http.MethodGet, "/palette?q=dash", nil)
	serve(w, q)
	assert.Integer(t, http.StatusOK, w.Code)
	assert.StringContains(t, "<strong>Dash</strong>board", w.Body.String())
}

func TestServeNoResults(t *testing.T) {
	serve := palette.NewServe(testRegistry())
	w := httptest.NewRecorder()
	q := httptest.NewRequest(http.MethodGet, "/palette?q=xyz", nil)
	serve(w, q)
	assert.Integer(t, http.StatusOK, w.Code)
	assert.StringContains(t, "No matches", w.Body.String())
}

func TestServeAcronymHighlight(t *testing.T) {
	serve := palette.NewServe(testRegistry())
	w := httptest.NewRecorder()
	q := httptest.NewRequest(http.MethodGet, "/palette?q=cp", nil)
	serve(w, q)
	assert.Integer(t, http.StatusOK, w.Code)
	assert.StringContains(t, "<strong>C</strong>reate", w.Body.String())
	assert.StringContains(t, "<strong>p</strong>roject", w.Body.String())
}

func TestServeResultsAreLinks(t *testing.T) {
	serve := palette.NewServe(testRegistry())
	w := httptest.NewRecorder()
	q := httptest.NewRequest(http.MethodGet, "/palette?q=metric", nil)
	serve(w, q)
	assert.Integer(t, http.StatusOK, w.Code)
	assert.StringContains(t, "/metrics", w.Body.String())
}
