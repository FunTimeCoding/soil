package service

func (s *Service) StopContainer(
	instance string,
	identifier int,
	node string,
) (string, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return "", clientFail
	}

	container, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.StopContainer(container)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
