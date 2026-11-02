package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"net/http"
	"sync/atomic"
	"testing"
)

func TestNotifierNotify(t *testing.T) {
	sink, c := worker_tester.NewSink(t)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.NewRosterGuard(t, c, r))
	assert.FatalOnError(t, n.Notify("kilo", "first message"))
	received := sink.All()
	assert.Integer(t, 1, len(received))
	assert.String(t, "kilo", received[0].Callsign)
	assert.String(t, "mattermost", received[0].Source)
	assert.String(t, "first message", received[0].Body)
	assert.Integer(t, 0, len(r.Events()))
}

func TestNotifierNotifySkipsEmptyCallsign(t *testing.T) {
	var calls atomic.Int32
	c := answering(t, http.StatusOK, &calls)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.NewRosterGuard(t, c, r))
	assert.Nil(t, n.Notify("", "first message"))
	assert.Integer(t, 0, int(calls.Load()))
}

func TestNotifierAnswersAFailureForItsCallerToReport(t *testing.T) {
	var calls atomic.Int32
	c := answering(t, http.StatusInternalServerError, &calls)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.NewRosterGuard(t, c, r))
	e := n.Notify("kilo", "first message")
	assert.True(t, classify.Reportable(e))
	assert.Integer(t, 0, len(r.Events()))
}

func TestAGoclaudedOutageHoldsTheNextDigestOff(t *testing.T) {
	var calls atomic.Int32
	c := answering(t, http.StatusServiceUnavailable, &calls)
	r := memory.New()
	n := notifier.New(c, "mattermost", worker_tester.NewRosterGuard(t, c, r))
	assert.True(t, unavailable.Is(n.Notify("kilo", "first")))
	assert.True(t, unavailable.Is(n.Notify("kilo", "second")))
	assert.Integer(t, 1, int(calls.Load()))
	assert.Integer(t, 0, len(r.Events()))
}
