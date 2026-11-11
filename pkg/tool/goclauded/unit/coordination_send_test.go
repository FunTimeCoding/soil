package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestSendDirectMessage(t *testing.T) {
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
			constant.Body: "heads up: running lint",
		},
	)
	r := b.CheckLive()
	messages := clientKindEntries(r.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
	assert.StringContains(t, a.Name(), messages[0].Body)
	assert.StringContains(t, "heads up: running lint", messages[0].Body)
}

func TestImmediateBroadcastSaysItWasQueued(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	answer := a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.Body:              "foxtrot",
			constant.DeliveryParameter: "immediate",
		},
	)
	assert.StringContains(t, "queued - a broadcast never wakes anyone", answer)
}

func TestRestReportsWhetherSendWokeAnyone(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	broadcast, e := a.RestClient.PostSendWithResponse(
		a.Context,
		client.SendRequest{
			Callsign: a.Name(),
			Body:     "golf",
			Delivery: new(client.Immediate),
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, broadcast.JSON200.Immediate)
	named, f := a.RestClient.PostSendWithResponse(
		a.Context,
		client.SendRequest{
			Callsign: a.Name(),
			Body:     "hotel",
			Delivery: new(client.Immediate),
			To:       &[]string{b.Name()},
		},
	)
	assert.FatalOnError(t, f)
	assert.True(t, named.JSON200.Immediate)
}

func TestRestSendPastImmediateLimitSaysItWasHeld(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	request := client.SendRequest{
		Callsign: a.Name(),
		Body:     "tango",
		Delivery: new(client.Immediate),
		To:       &[]string{b.Name()},
	}

	for range constant.ImmediateLimit {
		r, e := a.RestClient.PostSendWithResponse(a.Context, request)
		assert.FatalOnError(t, e)
		assert.True(t, r.JSON200.Immediate)
	}

	r, e := a.RestClient.PostSendWithResponse(a.Context, request)
	assert.FatalOnError(t, e)
	assert.False(t, r.JSON200.Immediate)
}

func TestBroadcastReachesAnotherSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "listener")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{constant.Body: "deploying service"},
	)
	r := b.CheckLive()
	messages := clientKindEntries(r.Entries, constant.QueueMessage)
	assert.Count(t, 1, messages)
	assert.StringContains(t, "deploying service", messages[0].Body)
}
