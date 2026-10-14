package reporter

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/getsentry/sentry-go"
)

type Reporter struct {
	hub     *sentry.Hub
	client  *sentry.Client
	project string
	locator string
	version string
	pool    *pool.Pool
}
