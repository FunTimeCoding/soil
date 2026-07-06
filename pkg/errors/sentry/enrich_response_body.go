package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/getsentry/sentry-go"
)

func enrichResponseBody(
	e *sentry.Event,
	h *sentry.EventHint,
) *sentry.Event {
	if h == nil {
		return e
	}

	if b := extractBody(h.OriginalException); b != nil {
		e.Contexts[constant.Response] = sentry.Context{
			constant.Body: string(b),
		}

		return e
	}

	if f, okay := h.RecoveredException.(error); okay {
		if b := extractBody(f); b != nil {
			e.Contexts[constant.Response] = sentry.Context{
				constant.Body: string(b),
			}
		}
	}

	return e
}
