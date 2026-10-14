package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"testing"
)

func TestClientWalksTheCursorFromTheHead(t *testing.T) {
	s := base.New(t)
	c := s.AlertClient(t)
	head, e := c.LatestIdentifier()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, int(head))
	all, f := c.AlertsAfter(0, 5)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, all)
	assert.String(t, "HighMemory", all[0].Name)
	assert.True(t, all[0].Firing())
	after, g := c.AlertsAfter(head, 5)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, after)
}

func TestClientSeesAResolutionByIdentifier(t *testing.T) {
	s := base.New(t)
	c := s.AlertClient(t)
	s.MockClient.Remove("fp1")
	s.Worker.Poll()
	result, e := c.AlertsByIdentifier([]uint{1, 2})
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
	assert.False(t, result[0].Firing())
	assert.True(t, result[1].Firing())
}
