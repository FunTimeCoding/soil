package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/getsentry/sentry-go"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnIncomingTraceIsOnTheRequestHub(t *testing.T) {
	p := serveTrace(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"GET /trace",
				func(
					w http.ResponseWriter,
					q *http.Request,
				) {
					if h := sentry.GetHubFromContext(q.Context()); h != nil {
						writer.Print(w, "%s", h.GetTraceparent())
					}
				},
			)
		},
	)
	locator := fmt.Sprintf("http://localhost:%d/trace", p)
	continued := tracedGet(
		t,
		locator,
		"0123456789abcdef0123456789abcdef-0123456789abcdef-1",
	)
	assert.Prefix(t, "0123456789abcdef0123456789abcdef-", continued)
	fresh := tracedGet(t, locator, "")
	assert.True(t, fresh != "")
	assert.StringNotContains(t, "0123456789abcdef0123456789abcdef", fresh)
}

func TestAnOutgoingRequestCarriesTheIncomingTrace(t *testing.T) {
	var seen string
	d := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				seen = q.Header.Get(sentry.SentryTraceHeader)
			},
		),
	)
	defer d.Close()
	o := stall_client.New().Traced().Build()
	defer o.Close()
	p := serveTrace(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"GET /call",
				func(
					_ http.ResponseWriter,
					q *http.Request,
				) {
					r, e := o.Client().Do(
						web.NewGet(d.URL).WithContext(q.Context()),
					)
					errors.PanicOnError(e)
					errors.PanicClose(r.Body)
				},
			)
		},
	)
	tracedGet(
		t,
		fmt.Sprintf("http://localhost:%d/call", p),
		"0123456789abcdef0123456789abcdef-0123456789abcdef-1",
	)
	assert.Prefix(t, "0123456789abcdef0123456789abcdef-", seen)
}
