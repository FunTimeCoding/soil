package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/node_status"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

func (s *Service) GetNodeStatus(
	c face.ProxmoxClient,
	node string,
) (*node_status.Status, error) {
	return c.NodeStatus(node)
}
