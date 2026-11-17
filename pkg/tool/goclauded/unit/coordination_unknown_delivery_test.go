package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/delivery_tester"
	"testing"
)

func TestRestSendRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	b := s.NewSession()
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	r, e := a.GeneratedClient.PostSendWithResponse(
		a.Context,
		client.SendRequest{
			Callsign: a.Name(),
			Body:     "quebec",
			To:       &[]string{b.Name()},
			Delivery: new(client.Delivery(constant.FixtureUnknownDelivery)),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.Count(
		t,
		0,
		delivery_tester.ClientKindEntries(
			b.CheckLive().Entries,
			constant.QueueMessage,
		),
	)
}

func TestRestNotifyRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.GeneratedClient.PostNotifyWithResponse(
		a.Context,
		client.NotifyRequest{
			Callsign: a.Name(),
			Source:   "downloader",
			Body:     "romeo",
			Delivery: new(client.Delivery(constant.FixtureUnknownDelivery)),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.Count(
		t,
		0,
		delivery_tester.ClientKindEntries(
			a.CheckLive().Entries,
			constant.QueueNotification,
		),
	)
}

func TestRestPulseRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.GeneratedClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{
			Body:     "sierra",
			Delivery: new(client.Delivery(constant.FixtureUnknownDelivery)),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.Count(
		t,
		0,
		delivery_tester.ClientKindEntries(
			a.CheckLive().Entries,
			constant.QueuePulse,
		),
	)
}
