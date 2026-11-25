package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/funtimecoding/soil/pkg/watch/notifier"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestNotifierNotify(t *testing.T) {
	sink, c := worker_tester.New(t).Sink()
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.New(t).RosterGuard(c, r))
	assert.FatalOnError(
		t,
		n.Notify("kilo", "first message", goclaudedConstant.DeliveryImmediate),
	)
	received := sink.All()
	assert.Integer(t, 1, len(received))
	assert.String(t, "kilo", received[0].Callsign)
	assert.String(t, "mattermost", received[0].Source)
	assert.String(t, "first message", received[0].Body)
	assert.String(t, "immediate", string(*received[0].Delivery))
	assert.Integer(t, 0, len(r.Events()))
}

func TestNotifierNotifySkipsEmptyCallsign(t *testing.T) {
	s := upstream_server.New(t)
	c := client.New(s.URL, false, "")
	t.Cleanup(c.Close)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.New(t).RosterGuard(c, r))
	assert.Nil(
		t,
		n.Notify("", "first message", goclaudedConstant.DeliveryImmediate),
	)
	assert.Integer(t, 0, s.Count())
}

func TestNotifierAnswersFailureForItsCallerToReport(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
	)
	c := client.New(s.URL, false, "")
	t.Cleanup(c.Close)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.New(t).RosterGuard(c, r))
	e := n.Notify("kilo", "first message", goclaudedConstant.DeliveryImmediate)
	assert.True(t, classify.Reportable(e))
	assert.Integer(t, 0, len(r.Events()))
}

func TestGoclaudedOutageHoldsNextDigestOff(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
	)
	c := client.New(s.URL, false, "")
	t.Cleanup(c.Close)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.New(t).RosterGuard(c, r))
	assert.True(
		t,
		unavailable.Is(
			n.Notify("kilo", "first", goclaudedConstant.DeliveryImmediate),
		),
	)
	assert.True(
		t,
		unavailable.Is(
			n.Notify("kilo", "second", goclaudedConstant.DeliveryImmediate),
		),
	)
	assert.Integer(t, 1, s.Count())
	assert.Integer(t, 0, len(r.Events()))
}
