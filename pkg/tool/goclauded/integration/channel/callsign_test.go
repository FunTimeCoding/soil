package channel

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/channel_tester"
	"net/http"
	"testing"
	"time"
)

func TestChannelCallsignReturnsAfterAPromptTheChannelMissed(t *testing.T) {
	s := base.New(t)
	since := s.Store.Clock()().Add(-time.Second)
	r := s.Check("session-1")
	response, e := s.RESTClient(t).GetChannelCallsignWithResponse(
		t.Context(),
		&client.GetChannelCallsignParams{Session: "session-1", Since: since},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, response.StatusCode())
	assert.String(t, r.Callsign, response.JSON200.Callsign)
}

func TestChannelCallsignHoldsWhenTheOnlyPromptPredatesTheChannel(t *testing.T) {
	s := base.New(t)
	s.Check("session-1")
	s.Store.Advance(time.Minute)
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Callsign(t, s, x, s.Store.Clock()())

	select {
	case <-result:
		t.Fatal("callsign returned on a prompt older than the channel")
	case <-time.After(250 * time.Millisecond):
	}
}

func TestChannelCallsignReleasesOnTheNextPrompt(t *testing.T) {
	s := base.New(t)
	r := s.Check("session-1")
	s.Store.Advance(time.Minute)
	x, cancel := context.WithCancel(t.Context())
	defer cancel()
	result := channel_tester.Callsign(t, s, x, s.Store.Clock()())

	select {
	case <-result:
		t.Fatal("callsign returned before the next prompt")
	case <-time.After(250 * time.Millisecond):
	}

	s.Store.Advance(time.Minute)
	s.Check("session-1")

	select {
	case name := <-result:
		assert.String(t, r.Callsign, name)
	case <-time.After(2 * time.Second):
		t.Fatal("callsign did not arrive after the prompt")
	}
}
