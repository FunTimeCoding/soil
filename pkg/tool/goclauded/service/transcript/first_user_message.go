package transcript

func (s *Service) FirstUserMessage(sessionIdentifier string) string {
	return s.client.FirstUserMessage(sessionIdentifier)
}
