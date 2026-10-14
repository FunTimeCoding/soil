package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/getsentry/sentry-go"
	"testing"
)

func TestReporterClosesWithoutStart(t *testing.T) {
	t.Setenv(constant.LocatorEnvironment, "")
	r := reporter.NewOptional(stringsConstant.LowerAlfa)
	r.Close()
	r.Close()
	r.RecoverFlush(nil)
}

func TestReporterClosesTheClientItStarted(t *testing.T) {
	t.Setenv(constant.LocatorEnvironment, "https://key@sentry.example/1")
	t.Cleanup(func() { sentry.CurrentHub().BindClient(nil) })
	r := reporter.New(stringsConstant.LowerAlfa).Start()
	assert.NotNil(t, sentry.CurrentHub().Client())
	r.RecoverFlush(nil)
	r.Close()
}
