package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAQuietPanicRendersWithoutAReport(t *testing.T) {
	r := memory.New()
	recorder := httptest.NewRecorder()
	recoveryView().Recovery(r)(
		http.HandlerFunc(
			func(http.ResponseWriter, *http.Request) {
				panic(unavailable.Format("goalfa unreachable"))
			},
		),
	).ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, constant.RootPath, nil),
	)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	assert.StringContains(t, "goalfa unreachable", recorder.Body.String())
	assert.Count(t, 0, r.Events())
}

func TestAFaultingPanicIsReported(t *testing.T) {
	r := memory.New()
	recorder := httptest.NewRecorder()
	recoveryView().Recovery(r)(
		http.HandlerFunc(
			func(http.ResponseWriter, *http.Request) {
				panic(errors.New("bravo exploded"))
			},
		),
	).ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, constant.RootPath, nil),
	)
	assert.Count(t, 1, r.Events())
}

func TestAnAnswerCarriesTheStatusOfItsClass(t *testing.T) {
	r := memory.New()
	recorder := httptest.NewRecorder()
	recoveryView().Answer(
		recorder,
		httptest.NewRequest(http.MethodPost, constant.RootPath, nil),
		r,
		not_found.New("report", "charlie.html"),
	)
	assert.Integer(t, http.StatusNotFound, recorder.Code)
	assert.StringContains(t, "charlie.html", recorder.Body.String())
	assert.Count(t, 0, r.Events())
}

func TestAnAnswerToAFaultIsReported(t *testing.T) {
	r := memory.New()
	recorder := httptest.NewRecorder()
	recoveryView().Answer(
		recorder,
		httptest.NewRequest(http.MethodPost, constant.RootPath, nil),
		r,
		errors.New("delta exploded"),
	)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	assert.Count(t, 1, r.Events())
}
