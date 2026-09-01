package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
)

func (w *Worker) pollNode(
	c face.ProxmoxClient,
	hypervisor string,
	name string,
) *floor.Node {
	result := floor.Node{Hypervisor: hypervisor, Name: name}
	n, e := c.Node(name)

	if e != nil {
		w.log.Plain("node %s on %s unavailable: %v", name, hypervisor, e)

		return &result
	}

	if v, f := c.NodeVersion(n); f != nil {
		w.log.Plain("node %s version unavailable: %v", name, f)
	} else {
		w.collector.SetVersion(hypervisor, name, v)
		result.Version = v.Release
	}

	if u, f := c.UpdatesPending(n); f != nil {
		w.log.Plain("node %s pending updates unavailable: %v", name, f)
	} else {
		w.collector.SetUpdatePending(hypervisor, name, len(u))
		result.UpdatesPending = len(u)
	}

	return &result
}
