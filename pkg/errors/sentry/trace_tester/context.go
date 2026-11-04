package trace_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/getsentry/sentry-go"
	"testing"
)

func Context(
	t *testing.T,
	c *sentry.Client,
) context.Context {
	t.Helper()
	p, e := sentry.PropagationContextFromHeaders(
		"0123456789abcdef0123456789abcdef-0123456789abcdef-1",
		"sentry-trace_id=0123456789abcdef0123456789abcdef",
	)
	assert.FatalOnError(t, e)
	h := sentry.NewHub(c, sentry.NewScope())
	h.Scope().SetPropagationContext(p)

	return sentry.SetHubOnContext(t.Context(), h)
}
