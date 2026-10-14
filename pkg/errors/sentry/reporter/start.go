package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	sentryLibrary "github.com/getsentry/sentry-go"
)

func (r *Reporter) Start() *Reporter {
	if r.locator == "" {
		return r
	}

	r.client = sentry.Start(r.project, r.locator, r.version, r.pool.Client())
	r.hub = sentryLibrary.CurrentHub()

	return r
}
