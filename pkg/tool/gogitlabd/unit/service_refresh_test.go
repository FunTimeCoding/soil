package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestRefreshNotifiesOnlyWhenBoardMoves(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	c := o.Alfa()
	service_tester.SeedForge(c, 1, "app")
	n := s.Notifier().Subscribe()
	defer s.Notifier().Unsubscribe(n)
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.Signalled(t, n)
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.NotSignalled(t, n)
	c.SeedBranch("feature", "bbb")
	c.SeedPipeline(11, "feature", "bbb", "running")
	assert.Count(t, 0, s.Refresh(t.Context()))
	assert.Signalled(t, n)
	assert.Count(t, 2, s.Entries())
}
