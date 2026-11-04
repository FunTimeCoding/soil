package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/subscribed"
	"testing"
)

func TestAThreadDigestIsImmediateByDefault(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", "")
	o.Reply("bravo", "alfa", "first message")
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 1, o.Roster.Notes())
	assert.String(t, "immediate", string(o.Roster.Notes()[0].Delivery))
}

func TestAQueuedSubscriptionDigestIsQueued(t *testing.T) {
	o := subscribed.New(t)
	o.Subscribe(t, "alfa", "kilo", goclaudedConstant.DeliveryQueued)
	o.Reply("bravo", "alfa", "first message")
	assert.Count(t, 0, o.Service.FlushDue())
	assert.Count(t, 1, o.Roster.Notes())
	assert.String(t, "queued", string(o.Roster.Notes()[0].Delivery))
}
