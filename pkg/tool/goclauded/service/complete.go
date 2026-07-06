package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) Complete(
	sessionIdentifier string,
	name string,
	topic string,
	message string,
) error {
	c, e := s.store.UpsertCompletion(
		sessionIdentifier,
		name,
		constant.Complete,
		topic,
		message,
	)

	if e != nil {
		return e
	}

	if f := s.store.UpsertEvent(
		sessionIdentifier,
		constant.Complete,
		name,
		map[string]string{constant.Topic: topic, constant.Message: message},
	); f != nil {
		return f
	}

	if g := s.PushQueueBroadcast(
		constant.QueueSessionComplete,
		fmt.Sprintf("%s completed %s", name, topic),
	); g != nil {
		return g
	}

	slug, metadata, f := s.sessionMetadata(sessionIdentifier)

	if f != nil {
		return f
	}

	if slug == "" {
		return nil
	}

	metadata["created_at"] = c.CreatedAt.Format(time.RFC3339)

	return s.pushCompletion(slug, c.Sequence, message, metadata)
}
