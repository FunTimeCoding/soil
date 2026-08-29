package service

func (s *Service) RollbackMachineSnapshot(
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

	task, e := c.RollbackMachineSnapshot(vm, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
