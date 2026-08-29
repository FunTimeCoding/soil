package mock_service

import "github.com/funtimecoding/soil/pkg/errors/conflict"

func (s *Service) DeleteMachine(
	instance string,
	identifier int,
	node string,
	_ bool,
) error {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return clientFail
	}

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
