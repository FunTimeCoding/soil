package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/mattermost_server"
	"testing"
)

func TestSubscribeThreadStoresAndIndexes(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	result := c.MustCallToolNotation(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterAlias:    "papa",
		},
	)
	assert.Any(t, "alfa", result[constant.ParameterRoot])
	assert.Any(t, "bravo", result["channel"])
	stored := s.Store.MustByCallsign("kilo")
	assert.Integer(t, 1, len(stored))
	assert.String(t, "papa", stored[0].Alias)
	assert.String(t, "bravo", stored[0].ChannelIdentifier)
}

func TestSubscribeThreadAnswersLatestReply(t *testing.T) {
	s := base.New(t, mattermost_server.Replied, base.WithConnector())
	c := s.Tools()
	result := c.MustCallToolNotation(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
		},
	)
	latest, okay := result[constant.LatestReply].(map[string]any)
	assert.True(t, okay)
	assert.Any(t, "delta", latest["id"])
	assert.Any(t, "latest reply", latest["message"])
	assert.Any(t, "foxtrot", latest["username"])
}

func TestSubscribeThreadResolvesReplyToRoot(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	result := c.MustCallToolNotation(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "reply",
			constant.ParameterCallsign: "kilo",
		},
	)
	assert.Any(t, "alfa", result[constant.ParameterRoot])
	stored := s.Store.MustByCallsign("kilo")
	assert.String(t, "alfa", stored[0].RootIdentifier)
}

func TestSubscribeThreadIsIdempotent(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	arguments := map[string]any{
		constant.ParameterRoot:     "alfa",
		constant.ParameterCallsign: "kilo",
	}
	c.MustCallTool(constant.SubscribeThread, arguments)
	c.MustCallTool(constant.SubscribeThread, arguments)
	assert.Integer(t, 1, len(s.Store.MustByCallsign("kilo")))
}

func TestSubscribeThreadStoresQueuedDelivery(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	c.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterDelivery: "queued",
		},
	)
	assert.String(t, "queued", s.Store.MustByCallsign("kilo")[0].Delivery)
}

func TestSubscribeThreadRejectsUnknownDelivery(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	c.MustCallToolError(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterDelivery: "cruise",
		},
	)
	assert.Count(t, 0, s.Store.MustByCallsign("kilo"))
}

func TestSubscribeThreadRejectsMissingCallsign(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	assert.StringContains(
		t,
		"missing property 'callsign'",
		c.MustCallToolError(
			constant.SubscribeThread,
			map[string]any{constant.ParameterRoot: "alfa"},
		),
	)
}

func TestSubscribeThreadRejectsEmptyCallsign(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	assert.StringContains(
		t,
		"callsign is required",
		c.MustCallToolError(
			constant.SubscribeThread,
			map[string]any{
				constant.ParameterRoot:     "alfa",
				constant.ParameterCallsign: "",
			},
		),
	)
}

func TestUnsubscribeThreadRemovesEachSubscription(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	c.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
		},
	)
	c.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "lima",
		},
	)
	c.MustCallTool(
		constant.UnsubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
		},
	)
	assert.Integer(t, 0, len(s.Store.MustByCallsign("kilo")))
	assert.Integer(t, 1, len(s.Store.MustByCallsign("lima")))
	c.MustCallTool(
		constant.UnsubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "lima",
		},
	)
	assert.Integer(t, 0, len(s.Store.MustByCallsign("lima")))
}

func TestListSubscriptionsReturnsLabels(t *testing.T) {
	s := base.New(t, mattermost_server.Thread, base.WithConnector())
	c := s.Tools()
	c.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterAlias:    "papa",
		},
	)
	result := c.MustCallToolNotation(
		constant.ListSubscriptions,
		map[string]any{constant.ParameterCallsign: "kilo"},
	)
	rows, okay := result["subscriptions"].([]any)
	assert.True(t, okay)
	assert.Integer(t, 1, len(rows))
	row, rowOkay := rows[0].(map[string]any)
	assert.True(t, rowOkay)
	assert.Any(t, "papa", row[soil.LabelKey])
	assert.Any(t, "alfa", row[constant.ParameterRoot])
}
