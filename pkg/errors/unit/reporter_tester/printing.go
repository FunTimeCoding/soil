package reporter_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"io"
)

func (o *Tester) Printing(w io.Writer) *reporter.Reporter {
	o.t.Helper()

	return o.Sending("http://key@localhost:9/1", w)
}
