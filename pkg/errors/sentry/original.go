package sentry

import "github.com/getsentry/sentry-go"

func original(h *sentry.EventHint) error {
	if h.OriginalException != nil {
		return h.OriginalException
	}

	if f, okay := h.RecoveredException.(error); okay {
		return f
	}

	return nil
}
