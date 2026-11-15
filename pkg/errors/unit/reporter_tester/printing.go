package reporter_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"io"
	"testing"
)

func Printing(
	t *testing.T,
	w io.Writer,
) *reporter.Reporter {
	t.Helper()

	return Sending(t, "http://key@localhost:9/1", w)
}
