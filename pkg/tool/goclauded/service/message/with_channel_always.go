package message

func (s *Service) WithChannelAlways() *Service {
	s.channelAlways = true

	return s
}
