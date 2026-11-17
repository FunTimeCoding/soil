package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/mattermost_server"
	"testing"
)

func TestSubscriptionToolsAbsentWithoutGoclauded(t *testing.T) {
	name := base.New(t, mattermost_server.Thread).Tools().ToolNames()
	assert.True(t, len(name) > 0)

	for _, e := range name {
		assert.True(t, e != constant.SubscribeThread)
		assert.True(t, e != constant.UnsubscribeThread)
		assert.True(t, e != constant.ListSubscriptions)
	}
}

func TestSubscriptionToolsPresentWithGoclauded(t *testing.T) {
	name := base.New(
		t,
		mattermost_server.Thread,
		base.WithConnector(),
	).Tools().ToolNames()
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
