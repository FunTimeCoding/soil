package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCheckContextCutsALongMessage(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.To:   b.Name(),
			constant.Body: strings.Repeat("a", 5000),
		},
	)
	r := b.CheckLive()
	assert.StringContains(t, "…cut here - ", r.Context)
	assert.StringContains(t, "read it with read_message", r.Context)
	assert.True(t, utf8.RuneCountInString(r.Context) <= 2000)
}

func TestCheckPreviewRendersWithoutConsuming(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	b := s.NewSession(t)
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	b.CheckLive()
	a.MustCallTool(
		constant.Send,
		map[string]any{constant.To: b.Name(), constant.Body: "look at this"},
	)
	preview, e := b.RestClient.GetCheckWithResponse(
		b.Context,
		&client.GetCheckParams{Session: b.UUID, Preview: new(true)},
	)
	assert.FatalOnError(t, e)
	assert.StringContains(t, "look at this", preview.JSON200.Context)
	assert.StringContains(t, "look at this", b.CheckLive().Context)
}

func TestRestReadMessages(t *testing.T) {
	s := base.New(t)
	m := s.Store.SendMessage("Ash", "Cedar", "an answer")
	r, e := s.RESTClient(t).GetMessagesWithResponse(
		t.Context(),
		&client.GetMessagesParams{Identifier: []int{int(m.Identifier), 999}},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.Count(t, 1, r.JSON200.Messages)
	assert.String(t, "Ash", r.JSON200.Messages[0].From)
	assert.String(t, "Cedar", r.JSON200.Messages[0].To)
	assert.String(t, "an answer", r.JSON200.Messages[0].Body)
	assert.Count(t, 1, r.JSON200.Missing)
	assert.Integer(t, 999, r.JSON200.Missing[0])
}

func TestRestNotifyOverLimitIsRefused(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostNotifyWithResponse(
		a.Context,
		client.NotifyRequest{
			Callsign: a.Name(),
			Source:   "mattermost",
			Body:     strings.Repeat("b", 501),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.String(
		t,
		"notification is 501 characters; summarize it to 500 or fewer",
		r.JSON400.Error,
	)
}

func TestRestPulseOverLimitIsRefused(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{Body: strings.Repeat("c", 501)},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.String(
		t,
		"pulse is 501 characters; summarize it to 500 or fewer",
		r.JSON400.Error,
	)
}
