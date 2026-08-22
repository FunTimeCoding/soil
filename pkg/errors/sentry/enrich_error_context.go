package sentry

import "github.com/getsentry/sentry-go"

func enrichErrorContext(
	e *sentry.Event,
	h *sentry.EventHint,
) *sentry.Event {
	if h == nil {
		return e
	}

	if key, context := extractContext(h.OriginalException); context != nil {
		e.Contexts[key] = sentry.Context(context)

		return e
	}

	if f, okay := h.RecoveredException.(error); okay {
		if key, context := extractContext(f); context != nil {
			e.Contexts[key] = sentry.Context(context)
		}
	}

	return e
}
