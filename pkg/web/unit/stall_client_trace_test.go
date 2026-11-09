package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"testing"
)

func TestUntracedClientSendsNoTrace(t *testing.T) {
	trace, baggage := traceHeaders(t, stall_client.New().Build())
	assert.String(t, "", trace)
	assert.String(t, "", baggage)
}

func TestTracedClientSendsTrace(t *testing.T) {
	trace, baggage := traceHeaders(t, stall_client.New().Traced().Build())
	assert.Prefix(t, "0123456789abcdef0123456789abcdef-", trace)
	assert.StringContains(t, "sentry-release=alfa", baggage)
}
