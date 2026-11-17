package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/netbox/mock_client"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/http"
	"testing"
)

func TestNetBoxOutageIsAnsweredUnavailableWithoutCaptureOverRest(t *testing.T) {
	c := mock_client.New(t)
	c.Failure = outage_tester.RefusedDial()
	r := memory.New()
	s := base.GeneratedClient(t, base.New(t, c, r))

	for range 2 {
		answer, e := s.ListDevicesWithResponse(t.Context(), nil)
		assert.FatalOnError(t, e)
		assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode())
		assert.FatalNotNil(t, answer.JSON503)
	}

	assert.Count(t, 0, r.Events())
}

func TestNetBoxOutageFailsToolWithoutCapture(t *testing.T) {
	c := mock_client.New(t)
	c.Failure = outage_tester.RefusedDial()
	r := memory.New()
	m := base.New(t, c, r).Connect(t)
	defer m.Close()
	assert.StringContains(
		t,
		"connection refused",
		m.MustCallToolError(constant.ListDevices, nil),
	)
	assert.Count(t, 0, r.Events())
}

func TestNetBoxOutageAnswersBookmarkPageUnavailable(t *testing.T) {
	c := mock_client.New(t)
	c.Failure = outage_tester.RefusedDial()
	r := memory.New()
	s := base.New(t, c, r)
	answer, e := s.Client().Get(s.Address(constant.BookmarkPath))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(answer.Body)
	assert.Integer(t, http.StatusServiceUnavailable, answer.StatusCode)
	assert.Count(t, 0, r.Events())
}

func TestDaemonHandsOutNetBoxGuardFirst(t *testing.T) {
	d := service_tester.New(t, service_tester.WithPowerIntent()).Daemon()
	_, okay := d.Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestWaitAnswersOnceNetBoxAnswers(t *testing.T) {
	assert.FatalOnError(t, service_tester.New(t).Daemon().Wait(t.Context()))
}
