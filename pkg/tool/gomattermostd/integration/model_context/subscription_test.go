package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/model_context_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
	"testing"
)

func upstream(m *http.ServeMux) {
	m.HandleFunc(
		"/api/v4/posts/alfa",
		func(w http.ResponseWriter, _ *http.Request) {
			web.Encode(w, &model.Post{Id: "alfa", ChannelId: "bravo"})
		},
	)
	m.HandleFunc(
		"/api/v4/posts/reply",
		func(w http.ResponseWriter, _ *http.Request) {
			web.Encode(
				w,
				&model.Post{
					Id:        "reply",
					RootId:    "alfa",
					ChannelId: "bravo",
				},
			)
		},
	)
}

func TestSubscribeThreadStoresAndIndexes(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	result := r.Client.MustCallToolNotation(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterAlias:    "papa",
		},
	)
	assert.Any(t, "alfa", result[constant.ParameterRoot])
	assert.Any(t, "bravo", result["channel"])
	stored := r.Server.Store.MustByCallsign("kilo")
	assert.Integer(t, 1, len(stored))
	assert.String(t, "papa", stored[0].Alias)
	assert.String(t, "bravo", stored[0].ChannelIdentifier)
	assert.Strings(t, []string{"alfa"}, r.Server.Indexer.Indexed())
}

func TestSubscribeThreadResolvesReplyToRoot(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	result := r.Client.MustCallToolNotation(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "reply",
			constant.ParameterCallsign: "kilo",
		},
	)
	assert.Any(t, "alfa", result[constant.ParameterRoot])
	assert.String(
		t,
		"alfa",
		r.Server.Store.MustByCallsign("kilo")[0].RootIdentifier,
	)
}

func TestSubscribeThreadIsIdempotent(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	arguments := map[string]any{
		constant.ParameterRoot:     "alfa",
		constant.ParameterCallsign: "kilo",
	}
	r.Client.MustCallTool(constant.SubscribeThread, arguments)
	r.Client.MustCallTool(constant.SubscribeThread, arguments)
	assert.Integer(t, 1, len(r.Server.Store.MustByCallsign("kilo")))
}

func TestSubscribeThreadRejectsMissingCallsign(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	assert.StringContains(
		t,
		"Missing:[callsign]",
		r.Client.MustCallToolError(
			constant.SubscribeThread,
			map[string]any{constant.ParameterRoot: "alfa"},
		),
	)
}

func TestSubscribeThreadRejectsEmptyCallsign(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	assert.StringContains(
		t,
		"callsign is required",
		r.Client.MustCallToolError(
			constant.SubscribeThread,
			map[string]any{
				constant.ParameterRoot:     "alfa",
				constant.ParameterCallsign: "",
			},
		),
	)
}

func TestUnsubscribeThreadForgetsWhenLastSubscriberLeaves(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	r.Client.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
		},
	)
	r.Client.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "lima",
		},
	)
	r.Client.MustCallTool(
		constant.UnsubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
		},
	)
	assert.Integer(t, 0, len(r.Server.Indexer.Forgotten()))
	r.Client.MustCallTool(
		constant.UnsubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "lima",
		},
	)
	assert.Strings(t, []string{"alfa"}, r.Server.Indexer.Forgotten())
	assert.Integer(t, 0, len(r.Server.Store.MustAll()))
}

func TestListSubscriptionsReturnsLabels(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	r.Client.MustCallTool(
		constant.SubscribeThread,
		map[string]any{
			constant.ParameterRoot:     "alfa",
			constant.ParameterCallsign: "kilo",
			constant.ParameterAlias:    "papa",
		},
	)
	result := r.Client.MustCallToolNotation(
		constant.ListSubscriptions,
		map[string]any{constant.ParameterCallsign: "kilo"},
	)
	rows, okay := result["subscriptions"].([]any)
	assert.True(t, okay)
	assert.Integer(t, 1, len(rows))
	row, rowOkay := rows[0].(map[string]any)
	assert.True(t, rowOkay)
	assert.Any(t, "papa", row[libraryConstant.LabelKey])
	assert.Any(t, "alfa", row[constant.ParameterRoot])
}
