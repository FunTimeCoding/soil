package sentry

import "github.com/getsentry/sentry-go"

func enrich(
	e *sentry.Event,
	h *sentry.EventHint,
) *sentry.Event {
	return enrichErrorContext(enrichResponseBody(e, h), h)
}
