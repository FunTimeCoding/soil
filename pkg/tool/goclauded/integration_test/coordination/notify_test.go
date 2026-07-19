package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestNotifyReachesSessionViaQueue(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostNotifyWithResponse(
		a.Context,
		client.NotifyRequest{
			Callsign: a.Name(),
			Source:   "ci-runner",
			Body:     "build complete: main@abc123",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	check := a.CheckLive()
	notifications := clientEntriesByKind(
		check.Entries,
		constant.QueueNotification,
	)
	assert.Count(t, 1, notifications)
	assert.StringContains(t, "ci-runner", notifications[0].Body)
	assert.StringContains(t, "build complete", notifications[0].Body)
}

func TestNotifyConsumedOnDrain(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostNotifyWithResponse(
		a.Context,
		client.NotifyRequest{
			Callsign: a.Name(),
			Source:   "ci-runner",
			Body:     "job finished",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	first := a.CheckLive()
	notifications := clientEntriesByKind(
		first.Entries,
		constant.QueueNotification,
	)
	assert.Count(t, 1, notifications)
	second := a.CheckLive()
	notifications = clientEntriesByKind(
		second.Entries,
		constant.QueueNotification,
	)
	assert.Count(t, 0, notifications)
}
