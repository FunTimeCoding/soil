package service

func (s *Service) UpdateGroup(
	groupIdentifier int,
	title string,
	color string,
) error {
	return s.client.UpdateGroup(groupIdentifier, title, color, nil)
}
