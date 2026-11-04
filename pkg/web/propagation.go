package web

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func propagation(q *http.Request) sentry.PropagationContext {
	result, e := sentry.PropagationContextFromHeaders(
		q.Header.Get(sentry.SentryTraceHeader),
		q.Header.Get(sentry.SentryBaggageHeader),
	)

	if e != nil {
		return sentry.NewPropagationContext()
	}

	return result
}
