package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestARecoveredNotFoundAnswersItsStatus(t *testing.T) {
	r := memory.New()
	wrapped := recoveryView().Recovery(r)(
		http.HandlerFunc(
			func(http.ResponseWriter, *http.Request) {
				panic(not_found.Format("paper alfa"))
			},
		),
	)
	recorder := httptest.NewRecorder()
	wrapped.ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, constant.RootPath, nil),
	)
	assert.Integer(t, http.StatusNotFound, recorder.Code)
	assert.StringContains(t, "paper alfa", recorder.Body.String())
	assert.StringContains(t, "<nav", recorder.Body.String())
	assert.Count(t, 0, r.Events())
}

func TestARecoveredConflictFragmentKeepsItsHeader(t *testing.T) {
	wrapped := recoveryView().Recovery(memory.New())(
		http.HandlerFunc(
			func(http.ResponseWriter, *http.Request) {
				panic(conflict.Format("already accepted"))
			},
		),
	)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, constant.RootPath, nil)
	request.Header.Set(constant.ExtendedRequest, "true")
	wrapped.ServeHTTP(recorder, request)
	assert.Integer(t, http.StatusConflict, recorder.Code)
	assert.String(t, "true", recorder.Header().Get(constant.NotificationItem))
}

func TestARecoveredUnclassifiedErrorStaysAFaultAndIsCaptured(t *testing.T) {
	r := memory.New()
	wrapped := recoveryView().Recovery(r)(
		http.HandlerFunc(
			func(http.ResponseWriter, *http.Request) {
				panic(errors.New("store failed"))
			},
		),
	)
	recorder := httptest.NewRecorder()
	wrapped.ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, constant.RootPath, nil),
	)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	assert.Count(t, 1, r.Events())
}
