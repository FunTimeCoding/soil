package unit

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"github.com/getsentry/sentry-go"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestAPrintingReporterStillSendsToSentry(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, "{}")
	var b bytes.Buffer
	r := sendingReporter(
		t,
		join.Empty("http://key@", strings.TrimPrefix(s.URL, "http://"), "/1"),
		&b,
	)
	assert.NotNil(t, r.Hub())
	r.CaptureException(fmt.Errorf("embed refused"))
	assert.True(t, r.Hub().Flush(2*time.Second))
	assert.NotNil(t, last())
	assert.StringContains(t, "/api/1/envelope/", last().URL.Path)
	assert.StringContains(t, "alfa: embed refused", b.String())
}

func TestAPrintingReporterPrintsACapturedException(t *testing.T) {
	var b bytes.Buffer
	r := printingReporter(t, &b)
	r.CaptureException(fmt.Errorf("embed refused"))
	assert.String(
		t,
		"alfa: embed refused\n",
		strings.SplitAfter(b.String(), "\n")[0],
	)
	assert.StringContains(
		t,
		"  pkg/errors/unit/reporter_print_test.go:",
		b.String(),
	)
	assert.StringContains(
		t,
		"TestAPrintingReporterPrintsACapturedException",
		b.String(),
	)
	assert.StringContains(t, "· 2 more frames: testing, runtime\n", b.String())
	assert.False(t, strings.Contains(b.String(), "sentry/reporter"))
}

func TestAPrintingReporterPrintsTheCaptureContext(t *testing.T) {
	var b bytes.Buffer
	r := printingReporter(t, &b)
	r.CaptureWithContext(
		fmt.Errorf("embed refused"),
		"request",
		map[string]any{"model": "bravo"},
	)
	assert.StringContains(t, "alfa: embed refused", b.String())
	assert.StringContains(t, "  request model bravo\n", b.String())
}

func TestAPrintingReporterPrintsARecoveredPanic(t *testing.T) {
	var b bytes.Buffer
	r := printingReporter(t, &b)
	r.Recover("charlie broke")
	assert.StringContains(t, "alfa: charlie broke", b.String())
	assert.StringContains(
		t,
		"TestAPrintingReporterPrintsARecoveredPanic",
		b.String(),
	)
}

func TestAReporterWithoutPrintPrintsNothing(t *testing.T) {
	t.Setenv(constant.LocatorEnvironment, "http://key@localhost:9/1")
	t.Cleanup(func() { sentry.CurrentHub().BindClient(nil) })
	var b bytes.Buffer
	r := reporter.New(stringsConstant.LowerAlfa).WithOutput(&b).Start()
	defer r.Close()
	r.CaptureException(fmt.Errorf("embed refused"))
	r.Recover("charlie broke")
	assert.String(t, "", b.String())
}
