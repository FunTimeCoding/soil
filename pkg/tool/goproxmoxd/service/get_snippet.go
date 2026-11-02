package service

func (s *Service) GetSnippet(
	instance string,
	name string,
) ([]byte, error) {
	path, e := snippetPath(name)

	if e != nil {
		return nil, e
	}

	c, e := s.shells.Client(instance)

	if e != nil {
		return nil, e
	}

	b, e := c.ReadFile(path)

	return b, s.forgetDropped(instance, c, e)
}
