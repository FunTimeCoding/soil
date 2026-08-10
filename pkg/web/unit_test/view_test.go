package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
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

func recoveryView() *view.View {
	return view.New(layout.New(identity.New("test", "test tool", "test")))
}

func TestRecoveryPage(t *testing.T) {
	v := recoveryView()
	wrapped := v.Recovery(memory.New())(http.HandlerFunc(panicServe))
	recorder := httptest.NewRecorder()
	wrapped.ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, constant.LocationRoot, nil),
	)
	assert.Integer(t, http.StatusInternalServerError, recorder.Code)
	body := recorder.Body.String()
	assert.True(t, strings.Contains(body, "notification-error"))
	assert.True(t, strings.Contains(body, "store failed"))
	assert.True(t, strings.Contains(body, "<nav"))
}

func renderBrand(l *layout.Page) string {
	recorder := httptest.NewRecorder()
	view.New(l).RenderPage(recorder, "", constant.LocationRoot)

	return recorder.Body.String()
}

func TestBrandLinksHome(t *testing.T) {
	body := renderBrand(layout.New(identity.New("test", "test tool", "test")))
	assert.True(t, strings.Contains(body, `<a href="/"><strong>test</strong></a>`))
	assert.True(t, !strings.Contains(body, "connection-dot"))
}

func TestLiveBrandCarriesConnectionDot(t *testing.T) {
	body := renderBrand(
		layout.New(identity.New("test", "test tool", "test")).
			WithLiveEndpoint("/event"),
	)
	assert.True(
		t,
		strings.Contains(
			body,
			`<a href="/"><span id="connection" class="connection-dot disconnected">`,
		),
	)
	assert.True(t, strings.Contains(body, "htmx:sseOpen"))
}

func TestRenderPageWithSummary(t *testing.T) {
	recorder := httptest.NewRecorder()
	recoveryView().RenderPageWithSummary(
		recorder,
		"Dashboard",
		constant.LocationRoot,
		[]string{"24 services", "2 drift"},
		html.H3(gomponents.Text("Fleet")),
	)
	body := recorder.Body.String()
	assert.True(t, strings.Contains(body, "24 services · 2 drift"))
	assert.True(
		t,
		strings.Index(body, "24 services") < strings.Index(body, "<h3>"),
	)
}

func TestRecoveryFragment(t *testing.T) {
	v := recoveryView()
	wrapped := v.Recovery(memory.New())(http.HandlerFunc(panicServe))
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, constant.LocationRoot, nil)
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
