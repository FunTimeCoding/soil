package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotifierNotify(t *testing.T) {
	sink, c := newNotifyClient(t)
	r := memory.New()
	notifier.New(c, "mattermost", r).Notify("kilo", "first message")
	received := sink.all()
	assert.Integer(t, 1, len(received))
	assert.String(t, "kilo", received[0].Callsign)
	assert.String(t, "mattermost", received[0].Source)
	assert.String(t, "first message", received[0].Body)
	assert.Integer(t, 0, len(r.Events()))
}

func TestNotifierNotifySkipsEmptyCallsign(t *testing.T) {
	count := 0
	s := httptest.NewServer(
		http.HandlerFunc(
			func(_ http.ResponseWriter, _ *http.Request) { count++ },
		),
	)
	defer s.Close()
	c := connector.New(s.URL, false, "")
	r := memory.New()
	notifier.New(c, "mattermost", r).Notify("", "first message")
	assert.Integer(t, 0, count)
	assert.Integer(t, 0, len(r.Events()))
}

func TestNotifierNotifyCapturesFailureStatus(t *testing.T) {
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
		),
	)
	defer s.Close()
	c := connector.New(s.URL, false, "")
	r := memory.New()
	notifier.New(c, "mattermost", r).Notify("kilo", "first message")
	assert.Integer(t, 1, len(r.Events()))
}
