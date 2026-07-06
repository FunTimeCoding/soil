package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func (s *Store) WaitMessage(
	name string,
	timeout time.Duration,
) ([]message.Message, error) {
	listening, e := s.isListening(name)

	if e != nil {
		return nil, e
	}

	if !listening {
		return nil, nil
	}

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		messages, e := s.PendingMessages(name)

		if e != nil {
			return nil, e
		}

		if len(messages) > 0 {
			return messages, nil
		}

		time.Sleep(2 * time.Second)
	}

	return nil, nil
}
