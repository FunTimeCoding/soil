package service

func (s *Service) DeleteBranch(listIdentifier int) error {
	links, e := s.LinksByList(listIdentifier)

	if e != nil {
		return e
	}

	for _, l := range links {
		if f := s.DeleteLink(l.Identifier); f != nil {
			return f
		}
	}

	return s.DeleteList(listIdentifier)
}
