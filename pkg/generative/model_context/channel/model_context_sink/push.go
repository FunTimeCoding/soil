package model_context_sink

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Sink) Push(
	content string,
	meta map[string]string,
) {
	s.server.SendNotificationToAllClients(
		constant.ChannelNotification,
		map[string]any{
			constant.ChannelContent: content,
			constant.ChannelMeta:    meta,
		},
	)
}
