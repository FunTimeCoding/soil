package base

func WithChannelAlways() Option {
	return func(s *setup) {
		s.channelAlways = true
	}
}
