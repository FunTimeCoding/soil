package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestRestSendRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	r, e := a.RestClient.PostSendWithResponse(
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
		clientKindEntries(b.CheckLive().Entries, constant.QueueMessage),
	)
}

func TestRestNotifyRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostNotifyWithResponse(
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
		clientKindEntries(a.CheckLive().Entries, constant.QueueNotification),
	)
}

func TestRestPulseRefusesUnknownDelivery(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
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
		clientKindEntries(a.CheckLive().Entries, constant.QueuePulse),
	)
}
