package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"net/http"
	"testing"
)

func TestNetBoxOutageIsAnsweredUnavailableWithoutCaptureOverRest(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	r := memory.New()
	s := newClient(t, newNetboxServer(t, c, r))

	for range 2 {
		answer, e := s.ListDevicesWithResponse(t.Context(), nil)
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, r.Events())
}

func TestNetBoxOutageFailsToolWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	r := memory.New()
	m := newNetboxServer(t, c, r).Connect(t)
	defer m.Close()
	assert.StringContains(
		t,
		"connection refused",
		m.MustCallToolError(constant.ListDevices, nil),
	)
	assert.Count(t, 0, r.Events())
}

func TestNetBoxOutageAnswersBookmarkPageUnavailable(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	r := memory.New()
	s := newNetboxServer(t, c, r)
	answer, e := s.Client().Get(s.Address(constant.BookmarkPath))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(answer.Body)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode)
	assert.Count(t, 0, r.Events())
}

func TestDaemonHandsOutNetBoxGuardFirst(t *testing.T) {
	o := service_tester.NewOption(t, service_tester.NewSource(t))
	o.PowerIntent = true
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
	t.Cleanup(d.Close)
	_, okay := d.Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceNetBoxAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		newDaemon(t, service_tester.NewSource(t)).Wait(t.Context()),
	)
}
