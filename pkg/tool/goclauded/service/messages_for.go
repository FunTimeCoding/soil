package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) messagesFor(
	entries []queue.Entry,
) (map[uint]*message.Message, error) {
	var identifiers []uint

	for _, e := range entries {
		if e.MessageIdentifier != nil {
			identifiers = append(identifiers, *e.MessageIdentifier)
		}
	}

	found, e := s.store.MessagesByIdentifiers(identifiers)

	if e != nil {
		return nil, e
	}

	result := map[uint]*message.Message{}

	for _, m := range found {
		result[m.Identifier] = m
	}

	return result, nil
}
