package service

func (s *Service) note(
	callsign string,
	e error,
) error {
	crossed, refused := s.outages.Note(callsign, e)

	if crossed != nil {
		s.logger.Plain("%s", crossed)
	}

	return refused
}
