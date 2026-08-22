package mock_service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
)

func (s *Service) DeleteMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
	_ bool,
) error {
	n, e := c.Node(node)

	if e != nil {
		return e
	}

	vm, e := c.Machine(n, identifier)

	if e != nil {
		return e
	}

	if vm.Status == "running" {
		return conflict.Format(
			"machine %d is running - stop it before deleting",
			identifier,
		)
	}

	return nil
}
