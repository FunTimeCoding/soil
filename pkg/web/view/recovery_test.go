package view

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func panicServe(
	http.ResponseWriter,
	*http.Request,
) {
	panic("store failed")
}

func recoveryView() *View {
	return New(layout.New(identity.New("test", "test tool", "test")))
}

func TestRecoveryPage(t *testing.T) {
	v := recoveryView()
	wrapped := v.Recovery(memory.New())(http.HandlerFunc(panicServe))
	recorder := httptest.NewRecorder()
	wrapped.ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, "/", nil),
	)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	body := recorder.Body.String()
	assert.True(t, strings.Contains(body, "notification-error"))
	assert.True(t, strings.Contains(body, "store failed"))
	assert.True(t, strings.Contains(body, "<nav"))
}

func TestRecoveryFragment(t *testing.T) {
	v := recoveryView()
	wrapped := v.Recovery(memory.New())(http.HandlerFunc(panicServe))
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set(constant.ExtendedRequest, "true")
	wrapped.ServeHTTP(recorder, request)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	assert.String(
		t,
		"true",
		recorder.Header().Get(constant.NotificationItem),
	)
	body := recorder.Body.String()
	assert.True(t, strings.Contains(body, "notification-error"))
	assert.True(t, !strings.Contains(body, "<nav"))
}
