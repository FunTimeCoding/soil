package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestRestSendWithoutDeliveryWakesNamedSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	r, e := a.RestClient.PostSendWithResponse(
		a.Context,
		client.SendRequest{
			Callsign: a.Name(),
			Body:     "kilo",
			To:       &[]string{b.Name()},
		},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.JSON200.Immediate)
}

func TestRestNotifyWithoutDeliveryStaysQueued(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostNotifyWithResponse(
		a.Context,
		client.NotifyRequest{
			Callsign: a.Name(),
			Source:   "mattermost",
			Body:     "lima",
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, r.JSON200.Immediate)
}

func TestRestPulseWithoutDeliveryStaysQueued(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{Body: "mike"},
	)
	assert.FatalOnError(t, e)
	assert.False(t, r.JSON200.Immediate)
}

func TestModelContextSendWithoutDeliveryWakesNamedSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.To:   []string{b.Name()},
			constant.Body: "november",
		},
	)
	r, e := b.RestClient.GetChannelWithResponse(
		b.Context,
		&client.GetChannelParams{Callsign: b.Name()},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.JSON200.Entries)
}

func TestModelContextBroadcastWithoutDeliverySaysItWasQueued(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	answer := a.MustCallTool(
		constant.Send,
		map[string]any{constant.Body: "oscar"},
	)
	assert.StringContains(t, "queued - a broadcast never wakes anyone", answer)
}

func TestModelContextSendRefusesCruise(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	answer := a.MustCallToolError(
		constant.Send,
		map[string]any{
			constant.To:                []string{b.Name()},
			constant.Body:              "papa",
			constant.DeliveryParameter: "cruise",
		},
	)
	assert.StringContains(
		t,
		"value must be one of 'immediate', 'queued'",
		answer,
	)
}

func TestDeliveryDescriptionWithoutChannelPromise(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	assert.StringContains(
		t,
		"immediate (default) wakes each named recipient now; one without a channel gets it with their next prompt.",
		sendToolNotation(t, a.ListTools()),
	)
}

func TestDeliveryDescriptionUnderChannelPromise(t *testing.T) {
	s := base.NewWithChannelAlways(t)
	a := s.NewSession(t)
	tool := sendToolNotation(t, a.ListTools())
	assert.StringContains(
		t,
		"immediate (default) wakes each named recipient now. queued arrives",
		tool,
	)
	assert.StringNotContains(t, "without a channel", tool)
}
