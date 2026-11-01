package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestConnectorSubscribeDeliversParsedEvents(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "connector topic")
	events, stop := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
		logger.New(context.Background()),
		memory.New(),
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
		logger.New(context.Background()),
		memory.New(),
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
		logger.New(context.Background()),
		memory.New(),
	)
	fixture.NextStreamEvent(t, first)
	stop()
	a.Announce(a.Name(), "later topic")
	second, halt := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
		logger.New(context.Background()),
		memory.New(),
	)
	defer halt()
	e := fixture.NextStreamEvent(t, second)
	assert.String(t, "later topic", e.Metadata[constant.Topic])
}

func TestConnectorSubscribeClosesItsChannelOnStop(t *testing.T) {
	s := base.New(t)
	events, stop := s.Connector(t).Subscribe(
		"consumer",
		[]string{constant.Announce},
		logger.New(context.Background()),
		memory.New(),
	)
	stop()

	for range events {
	}
}

func TestConnectorSubscribeReportsARefusedTokenOnce(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port).String(),
		false,
		"bogus",
	)
	defer c.Close()
	r := memory.New()
	_, stop := c.Subscribe(
		"consumer",
		[]string{constant.Announce},
		logger.New(context.Background()),
		r,
	)
	defer stop()
	awaitEvents(t, r, 1)
	assert.Count(t, 1, r.Events())
}
