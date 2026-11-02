package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestRefreshNotifiesOnlyWhenTheBoardMoves(t *testing.T) {
	s, clients := service_tester.New(t, constant.FixtureInstance)
	c := clients[constant.FixtureInstance]
	seedForge(c, 1, "app")
	n := s.Notifier().Subscribe()
	defer s.Notifier().Unsubscribe(n)
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.True(t, fired(n))
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.False(t, fired(n))
	c.SeedBranch("feature", "bbb")
	c.SeedPipeline(11, "feature", "bbb", "running")
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.True(t, fired(n))
	assert.Count(t, 2, s.Entries())
}
