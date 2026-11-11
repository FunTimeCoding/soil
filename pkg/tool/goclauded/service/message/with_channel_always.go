package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) WithChannelAlways() *Service {
	s.channelAlways = true
	s.defaultGrace = constant.DefaultCruiseGrace

	return s
}
