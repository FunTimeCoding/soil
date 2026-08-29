package service

func (s *Service) DeleteMachineSnapshot(
	instance string,
	identifier int,
	node string,
	name string,
) (string, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return "", clientFail
	}

	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.DeleteMachineSnapshot(vm, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
