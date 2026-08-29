package service

import "github.com/funtimecoding/soil/pkg/proxmox/node_status"

func (s *Service) GetNodeStatus(
	instance string,
	node string,
) (*node_status.Status, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	return c.NodeStatus(node)
}
