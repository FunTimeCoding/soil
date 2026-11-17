package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestRefreshFillsFloor(t *testing.T) {
	o := service_tester.New(t)
	o.Populate()
	s := o.Service()
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

func TestRefreshNotifiesOnlyWhenFloorChanges(t *testing.T) {
	o := service_tester.New(t)
	o.Populate()
	s := o.Service()
	n := s.Notifier().Subscribe()
	defer s.Notifier().Unsubscribe(n)
	s.Refresh()
	assert.Count(t, 1, n)
	<-n
	s.Refresh()
	assert.Count(t, 0, n)
	o.Alfa().Hypervisor.AddMachine("alfa", 102, "third")
	s.Refresh()
	assert.Count(t, 1, n)
}
