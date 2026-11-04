package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/getsentry/sentry-go"
	"io"
	"testing"
)

func tracedGet(
	t *testing.T,
	locator string,
	trace string,
) string {
	t.Helper()
	o := stall_client.New().Build()
	defer o.Close()
	q := web.NewGet(locator)

	if trace != "" {
		q.Header.Set(sentry.SentryTraceHeader, trace)
	}

	r, e := o.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	b, f := io.ReadAll(r.Body)
	assert.FatalOnError(t, f)

	return string(b)
}
