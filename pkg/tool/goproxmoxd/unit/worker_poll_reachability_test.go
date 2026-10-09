package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/worker_tester"
	"testing"
)

func TestPollTracksHypervisorReachability(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	w, _ := worker_tester.NewPollWorker("pve", c)
	w.Poll()
	assert.False(t, w.Unreachable("pve"))
	c.SetFailure(
		connection.New(constant.Unreachable, "pve", "", constant.Refused),
	)
	w.Poll()
	assert.True(t, w.Unreachable("pve"))
	w.Poll()
	assert.True(t, w.Unreachable("pve"))
	c.SetFailure(nil)
	w.Poll()
	assert.False(t, w.Unreachable("pve"))
}

func TestPollOtherFailureLeavesReachability(t *testing.T) {
	c := worker_tester.PopulatedClient("pve")
	w, _ := worker_tester.NewPollWorker("pve", c)
	w.Poll()
	c.SetFailure(errors.New("token rejected"))
	w.Poll()
	assert.False(t, w.Unreachable("pve"))
}
