package service

func (s *Service) StartMachine(
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

	task, e := c.StartMachine(vm)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
