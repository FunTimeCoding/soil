package sentry

import "github.com/getsentry/sentry-go"

func Enrich(
	e *sentry.Event,
	h *sentry.EventHint,
) *sentry.Event {
	return enrichConnection(enrichErrorContext(enrichResponseBody(e, h), h), h)
}
