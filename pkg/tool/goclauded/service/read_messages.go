package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"

func (s *Service) ReadMessages(identifiers []uint) ([]*message.Message, error) {
	return s.store.MessagesByIdentifiers(identifiers)
}
