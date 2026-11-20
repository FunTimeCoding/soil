package trace_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/getsentry/sentry-go"
	"io"
)

func (o *Tester) Get(
	locator string,
	trace string,
) string {
	o.t.Helper()
	c := stall_client.New().Build()
	defer c.Close()
	q := web.NewGet(locator)

	if trace != "" {
		q.Header.Set(sentry.SentryTraceHeader, trace)
	}

	r, e := c.Client().Do(q)
	assert.FatalOnError(o.t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(o.t, f)

	return string(b)
}
