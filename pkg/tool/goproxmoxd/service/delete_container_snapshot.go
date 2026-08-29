package service

func (s *Service) DeleteContainerSnapshot(
	instance string,
	identifier int,
	node string,
	name string,
) (string, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return "", clientFail
	}

	ct, e := findContainer(c, identifier, node)

	if e != nil {
		return "", e
	}

	task, e := c.DeleteContainerSnapshot(ct, name)

	if e != nil {
		return "", e
	}

	return string(task.UPID), nil
}
