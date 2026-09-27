package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"testing"
)

func TestConnectorSubscribeDeliversParsedEvents(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "connector topic")
	events, stop := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
	)
	defer stop()
	e := fixture.NextStreamEvent(t, events)
	assert.String(t, "announce", e.Kind)
	assert.String(t, "connector topic", e.Metadata[constant.Topic])
	assert.True(t, e.Identifier > 0)
}

func TestConnectorSubscribeFiltersKinds(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "ignored topic")
	a.MustCallTool(
		constant.Label,
		map[string]any{constant.Key: "colour", constant.Value: "green"},
	)
	events, stop := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Label},
	)
	defer stop()
	e := fixture.NextStreamEvent(t, events)
	assert.String(t, "label", e.Kind)
	assert.String(t, "green", e.Metadata[constant.Now])
}

func TestConnectorSubscribeResumesAfterStop(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "earlier topic")
	first, stop := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
	)
	fixture.NextStreamEvent(t, first)
	stop()
	a.Announce(a.Name(), "later topic")
	second, halt := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
	)
	defer halt()
	e := fixture.NextStreamEvent(t, second)
	assert.String(t, "later topic", e.Metadata[constant.Topic])
}
