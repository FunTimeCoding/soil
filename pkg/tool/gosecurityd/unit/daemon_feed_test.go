package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"github.com/funtimecoding/soil/pkg/worker"
	"testing"
)

func TestDaemonWithAnsibleProbesDaily(t *testing.T) {
	s := base.New(t)
	a := mock_client.New(t, constant.AnsiblePeer)
	o := s.Option()
	o.Ansible = func() face.Client { return a }
	d := daemon.New(o, logger.New(t.Context()), s.Reporter)
	workers := d.Workers()
	assert.Count(t, 7, workers)

	for _, w := range workers[:4] {
		_, guard := w.(*peer.Guard)
		assert.True(t, guard)
	}

	for _, w := range workers[4:] {
		_, clock := w.(*worker.Worker)
		assert.True(t, clock)
	}

	d.Close()
	assert.Integer(t, 1, a.Closed)
}
