package worker

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"

func (w *Worker) pollNode(
	c face.ProxmoxClient,
	hypervisor string,
	name string,
) {
	n, e := c.Node(name)

	if e != nil {
		w.log.Plain("node %s on %s unavailable: %v", name, hypervisor, e)

		return
	}

	if v, f := c.NodeVersion(n); f != nil {
		w.log.Plain("node %s version unavailable: %v", name, f)
	} else {
		w.collector.SetVersion(hypervisor, name, v)
	}

	if u, f := c.UpdatesPending(n); f != nil {
		w.log.Plain("node %s pending updates unavailable: %v", name, f)
	} else {
		w.collector.SetUpdatePending(hypervisor, name, len(u))
	}
}
