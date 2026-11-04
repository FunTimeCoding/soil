package trace_transport

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func (t *Transport) RoundTrip(e *http.Request) (*http.Response, error) {
	h := sentry.GetHubFromContext(e.Context())

	if h == nil {
		return t.transport.RoundTrip(e)
	}

	q := e.Clone(e.Context())
	q.Header.Set(sentry.SentryTraceHeader, h.GetTraceparent())

	if b := h.GetBaggage(); b != "" {
		q.Header.Set(sentry.SentryBaggageHeader, b)
	}

	return t.transport.RoundTrip(q)
}
