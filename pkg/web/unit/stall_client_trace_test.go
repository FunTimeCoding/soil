package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/trace_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"github.com/getsentry/sentry-go"
	"testing"
)

func TestUntracedClientSendsNoTrace(t *testing.T) {
	s := upstream_server.New(t)
	p := stall_client.New().Build()
	defer p.Close()
	r, e := p.Client().Do(
		web.NewGet(s.URL).WithContext(trace_tester.Context(t, nil)),
	)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)
	assert.String(t, "", s.Last().Header.Get(sentry.SentryTraceHeader))
	assert.String(t, "", s.Last().Header.Get(sentry.SentryBaggageHeader))
}

func TestTracedClientSendsTrace(t *testing.T) {
	s := upstream_server.New(t)
	p := stall_client.New().Traced().Build()
	defer p.Close()
	r, e := p.Client().Do(
		web.NewGet(s.URL).WithContext(trace_tester.Context(t, nil)),
	)
	assert.FatalOnError(t, e)
	errors.PanicClose(r.Body)
	assert.Prefix(
		t,
		"0123456789abcdef0123456789abcdef-",
		s.Last().Header.Get(sentry.SentryTraceHeader),
	)
	assert.StringContains(
		t,
		"sentry-trace_id=0123456789abcdef0123456789abcdef",
		s.Last().Header.Get(sentry.SentryBaggageHeader),
	)
}
