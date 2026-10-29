package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"io"
	"testing"
)

func printingReporter(
	t *testing.T,
	w io.Writer,
) *reporter.Reporter {
	t.Helper()

	return sendingReporter(t, "http://key@localhost:9/1", w)
}
