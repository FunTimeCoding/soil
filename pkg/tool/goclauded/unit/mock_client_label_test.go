package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"testing"
	"time"
)

func TestMockLeavesHeldLabelAloneAsClientDoes(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Join("alfa", 3, time.Now())
	claimed, e := c.ClaimLabel(
		"alfa",
		constant.FixtureLabelKey,
		"bravo",
		"operator",
	)
	assert.FatalOnError(t, e)
	assert.True(t, claimed)
	again, f := c.ClaimLabel(
		"alfa",
		constant.FixtureLabelKey,
		"watch",
		"reconciler",
	)
	assert.FatalOnError(t, f)
	assert.False(t, again)
}
