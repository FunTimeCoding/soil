package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"testing"
)

func TestARefreshFillsTheFloor(t *testing.T) {
	c := worker_tester.PopulatedClient("test")
	defer c.Close()
	s := service_tester.ContainerService(t, c)
	samples := s.Refresh()
	assert.Count(t, 1, samples)
	assert.FatalOnError(t, samples[0].Failure)
	f := s.Floor()
	assert.Count(t, 1, f.Nodes)
	assert.Integer(t, 1, f.Nodes[0].UpdatesPending)
	assert.Count(t, 2, f.Guests)
	assert.String(t, "first", f.Guests[0].Name)
	assert.False(t, f.Guests[0].Unbacked)
	assert.True(t, f.Guests[1].Unbacked)
}

func TestARefreshNotifiesOnlyWhenTheFloorChanges(t *testing.T) {
	c := worker_tester.PopulatedClient("test")
	defer c.Close()
	s := service_tester.ContainerService(t, c)
	n := s.Notifier().Subscribe()
	defer s.Notifier().Unsubscribe(n)
	s.Refresh()
	assert.Count(t, 1, n)
	<-n
	s.Refresh()
	assert.Count(t, 0, n)
	c.AddMachine("test", 102, "third")
	s.Refresh()
	assert.Count(t, 1, n)
}
