package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func (s *Service) AllMessages(
	limit int,
	skip int,
) []message.Message {
	return s.store.AllMessages(limit, skip)
}
