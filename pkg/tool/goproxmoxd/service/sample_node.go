package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/node_sample"
)

func (s *Service) sampleNode(
	c face.ProxmoxClient,
	hypervisor string,
	name string,
) *node_sample.NodeSample {
	result := node_sample.New(name)
	n, e := c.Node(name)

	if e != nil {
		s.logger.Plain("node %s on %s unavailable: %v", name, hypervisor, e)
		result.UpdatesFailure = e

		return result
	}

	if v, f := c.NodeVersion(n); f != nil {
		s.logger.Plain("node %s version unavailable: %v", name, f)
	} else {
		result.Version = v
	}

	result.Updates, result.UpdatesFailure = c.UpdatesPending(n)

	if result.UpdatesFailure != nil {
		s.logger.Plain(
			"node %s pending updates unavailable: %v",
			name,
			result.UpdatesFailure,
		)
	}

	return result
}
