package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/model_context_tester"
	"testing"
)

func TestSubscriptionToolsAbsentWithoutGoclauded(t *testing.T) {
	s := base.NewWithoutSubscription(t, threadUpstream)
	c := s.ContextServer.Connect(t)
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
	r := model_context_tester.New(t, threadUpstream)
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
