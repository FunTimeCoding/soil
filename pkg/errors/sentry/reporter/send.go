package reporter

import "github.com/getsentry/sentry-go"

func (r *Reporter) send(
	h *sentry.Hub,
	e error,
) string {
	if r.hub == nil {
		return ""
	}

	identifier := h.CaptureException(e)

	if identifier == nil {
		return ""
	}

	return string(*identifier)
}
