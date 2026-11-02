package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
)

func (s *Service) GetNodeStatus(
	instance string,
	node string,
) (*node_status.Status, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (*node_status.Status, error) {
			return c.NodeStatus(node)
		},
	)
}
