package reporter

import (
	"context"
	"github.com/getsentry/sentry-go"
)

func (r *Reporter) Capture(
	x context.Context,
	e error,
) string {
	if r.echo {
		r.print(e, "", nil)
	}

	h := sentry.GetHubFromContext(x)

	if h == nil {
		h = r.hub
	}

	return r.send(h, e)
}
