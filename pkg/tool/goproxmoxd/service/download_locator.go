package service

func (s *Service) DownloadLocator(
	instance string,
	node string,
	storage string,
	content string,
	filename string,
	l string,
) error {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return clientFail
	}

	n, e := c.Node(node)

	if e != nil {
		return e
	}

	st, e := c.Storage(n, storage)

	if e != nil {
		return e
	}

	task, e := c.DownloadLocator(st, content, filename, l)

	if e != nil {
		return e
	}

	return c.WaitForTask(task, 600)
}
