package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/model_context_tester"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"testing"
)

func toolNames(v []*mcp.Tool) []string {
	var result []string

	for _, e := range v {
		result = append(result, e.Name)
	}

	return result
}

func TestSubscriptionToolsAbsentWithoutGoclauded(t *testing.T) {
	s := base.NewWithoutSubscription(t, upstream)
	c := model_context_client.New(t, s.ContextServer.Port)
	t.Cleanup(
		func() {
			c.Close()
			s.Close()
		},
	)
	name := toolNames(c.ListTools())
	assert.True(t, len(name) > 0)

	for _, e := range name {
		assert.True(t, e != constant.SubscribeThread)
		assert.True(t, e != constant.UnsubscribeThread)
		assert.True(t, e != constant.ListSubscriptions)
	}
}

func TestSubscriptionToolsPresentWithGoclauded(t *testing.T) {
	r := model_context_tester.New(t, upstream)
	name := toolNames(r.Client.ListTools())
	found := 0

	for _, e := range name {
		if e == constant.SubscribeThread ||
			e == constant.UnsubscribeThread ||
			e == constant.ListSubscriptions {
			found++
		}
	}

	assert.Integer(t, 3, found)
}
