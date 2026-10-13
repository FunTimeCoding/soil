package reporter

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func (r *Reporter) Start() *Reporter {
	if r.locator == "" {
		return r
	}

	r.pool = stall_client.New().Build()
	r.hub = sentry.Start(r.project, r.locator, r.version, r.pool.Client())

	return r
}
