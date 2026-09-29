package service

func (s *Service) DeleteBranch(listIdentifier int) error {
	links, e := s.client.LinksByList(listIdentifier)

	if e != nil {
		return e
	}

	for _, l := range links {
		if f := s.client.DeleteLink(l.Identifier); f != nil {
			return f
		}
	}

	return s.client.DeleteList(listIdentifier)
}
