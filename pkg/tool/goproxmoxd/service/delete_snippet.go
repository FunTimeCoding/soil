package service

func (s *Service) DeleteSnippet(
	instance string,
	name string,
) error {
	path, e := snippetPath(name)

	if e != nil {
		return e
	}

	c, e := s.shells.Client(instance)

	if e != nil {
		return e
	}

	return s.forgetDropped(instance, c, c.RemoveFile(path))
}
