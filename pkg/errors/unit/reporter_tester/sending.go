package reporter_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/getsentry/sentry-go"
	"io"
)

func (o *Tester) Sending(
	locator string,
	w io.Writer,
) *reporter.Reporter {
	o.t.Helper()
	o.t.Setenv(constant.LocatorEnvironment, locator)
	o.t.Setenv(constant.PrintEnvironment, "1")
	o.t.Cleanup(func() { sentry.CurrentHub().BindClient(nil) })
	result := reporter.New(stringsConstant.LowerAlfa).WithOutput(w).Start()
	o.t.Cleanup(result.Close)

	return result
}
