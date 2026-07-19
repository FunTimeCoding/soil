package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestExternalPulseReachesSessionViaQueue(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{Body: "deploy complete"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	check := a.CheckLive()
	pulses := clientEntriesByKind(check.Entries, constant.QueuePulse)
	assert.Count(t, 1, pulses)
	assert.StringContains(t, "deploy complete", pulses[0].Body)
}

func TestMCPSelfPulseDoesNotEchoBack(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	a.MustCallTool(
		constant.Pulse,
		map[string]any{
			constant.Line: "still building",
		},
	)
	check := a.CheckLive()
	pulses := clientEntriesByKind(check.Entries, constant.QueuePulse)
	assert.Count(t, 0, pulses)
}

func TestExternalPulseConsumedOnDrain(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	r, e := a.RestClient.PostSessionPulseWithResponse(
		a.Context,
		a.UUID,
		client.PulseRequest{Body: "first pulse"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	first := a.CheckLive()
	pulses := clientEntriesByKind(first.Entries, constant.QueuePulse)
	assert.Count(t, 1, pulses)
	second := a.CheckLive()
	pulses = clientEntriesByKind(second.Entries, constant.QueuePulse)
	assert.Count(t, 0, pulses)
}
