package service

func (s *Service) CreateMachineSnapshot(
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

	task, e := c.CreateMachineSnapshot(vm, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
