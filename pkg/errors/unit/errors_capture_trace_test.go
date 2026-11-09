package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/sentry/trace_tester"
	"github.com/getsentry/sentry-go"
	"io"
	"testing"
)

func TestCaptureCarriesTraceOfItsContext(t *testing.T) {
	m := memory.New()
	classify.Capture(
		trace_tester.Context(t, nil),
		m,
		errors.New("disk full"),
		"",
	)
	assert.Count(t, 1, m.Events())
	assert.StringContains(
		t,
		"0123456789abcdef0123456789abcdef",
		m.Events()[0].Trace,
	)
}

func TestCaptureOutsideRequestHasNoTrace(t *testing.T) {
	m := memory.New()
	classify.Capture(t.Context(), m, errors.New("disk full"), "")
	assert.Count(t, 1, m.Events())
	assert.String(t, "", m.Events()[0].Trace)
}

func TestReporterSendsOnItsOwnHubOutsideRequest(t *testing.T) {
	r := sendingReporter(t, "http://key@localhost:1/1", io.Discard)
	m := &sentry.MockTransport{}
	c, e := sentry.NewClient(
		sentry.ClientOptions{Dsn: "http://key@localhost:1/1", Transport: m},
	)
	assert.FatalOnError(t, e)
	defer c.Close()
	sentry.CurrentHub().BindClient(c)
	r.Capture(t.Context(), errors.New("disk full"))
	assert.Count(t, 1, m.Events())
}

func TestReporterSendsOnHubOfItsContext(t *testing.T) {
	r := sendingReporter(t, "http://key@localhost:1/1", io.Discard)
	m := &sentry.MockTransport{}
	c, e := sentry.NewClient(
		sentry.ClientOptions{Dsn: "http://key@localhost:1/1", Transport: m},
	)
	assert.FatalOnError(t, e)
	defer c.Close()
	r.Capture(trace_tester.Context(t, c), errors.New("disk full"))
	assert.Count(t, 1, m.Events())
	assert.String(
		t,
		"0123456789abcdef0123456789abcdef",
		fmt.Sprint(m.Events()[0].Contexts["trace"]["trace_id"]),
	)
}
