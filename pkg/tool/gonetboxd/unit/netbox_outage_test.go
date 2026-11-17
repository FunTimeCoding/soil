package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestNetBoxOutageIsAnsweredUnavailableWithoutCaptureOverRest(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.GeneratedClient()

	for range 2 {
		answer, e := c.ListDevicesWithResponse(t.Context(), nil)
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestNetBoxOutageFailsToolWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	assert.StringContains(
		t,
		"connection refused",
		s.Tools().MustCallToolError(constant.ListDevices, nil),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestNetBoxOutageAnswersBookmarkPageUnavailable(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	v := s.Served()
	answer, e := v.Client().Get(v.Address(constant.BookmarkPath))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(answer.Body)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonHandsOutNetBoxGuardFirst(t *testing.T) {
	d := service_tester.New(t, service_tester.WithPowerIntent()).Daemon()
	_, okay := d.Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceNetBoxAnswers(t *testing.T) {
	assert.FatalOnError(t, service_tester.New(t).Daemon().Wait(t.Context()))
}
