package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"testing"
	"time"
)

func TestTheMockLeavesAHeldLabelAloneAsTheClientDoes(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Join("alfa", 3, time.Now())
	claimed, e := c.ClaimLabel("alfa", "role", "bravo", "operator")
	assert.FatalOnError(t, e)
	assert.True(t, claimed)
	again, f := c.ClaimLabel("alfa", "role", "watch", "reconciler")
	assert.FatalOnError(t, f)
	assert.False(t, again)
}
