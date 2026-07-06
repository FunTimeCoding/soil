package reporter

import "github.com/funtimecoding/soil/pkg/errors/sentry"

func (r *Reporter) Start() *Reporter {
	if r.locator == "" {
		return r
	}

	r.hub = sentry.Start(r.project, r.locator, r.version)

	return r
}
