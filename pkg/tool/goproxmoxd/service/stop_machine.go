package service

import netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"

func (s *Service) StopMachine(
	instance string,
	identifier int,
	node string,
) (string, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return "", clientFail
	}

	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.StopMachine(vm)

	if e != nil {
		return "", e
	}

	s.recordGuest(vm.Name, netbox.StatusOffline)

	return string(task.UPID), nil
}
