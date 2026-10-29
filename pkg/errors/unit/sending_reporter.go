package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/getsentry/sentry-go"
	"io"
	"testing"
)

func sendingReporter(
	t *testing.T,
	locator string,
	w io.Writer,
) *reporter.Reporter {
	t.Helper()
	t.Setenv(constant.LocatorEnvironment, locator)
	t.Setenv(constant.PrintEnvironment, "1")
	t.Cleanup(func() { sentry.CurrentHub().BindClient(nil) })
	result := reporter.New(stringsConstant.LowerAlfa).WithOutput(w).Start()
	t.Cleanup(result.Close)

	return result
}
