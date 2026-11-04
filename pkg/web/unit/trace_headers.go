package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/getsentry/sentry-go"
	"net/http"
	"net/http/httptest"
	"testing"
)

func traceHeaders(
	t *testing.T,
	p *pool.Pool,
) (string, string) {
	t.Helper()
	defer p.Close()
	var trace string
	var baggage string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				trace = q.Header.Get(sentry.SentryTraceHeader)
				baggage = q.Header.Get(sentry.SentryBaggageHeader)
			},
		),
	)
	defer s.Close()
	c, e := sentry.PropagationContextFromHeaders(
		"0123456789abcdef0123456789abcdef-0123456789abcdef-1",
		"sentry-trace_id=0123456789abcdef0123456789abcdef,sentry-release=alfa",
	)
	assert.FatalOnError(t, e)
	h := sentry.NewHub(nil, sentry.NewScope())
	h.Scope().SetPropagationContext(c)
	r, f := p.Client().Do(
		web.NewGet(s.URL).WithContext(sentry.SetHubOnContext(t.Context(), h)),
	)
	assert.FatalOnError(t, f)
	errors.PanicClose(r.Body)

	return trace, baggage
}
