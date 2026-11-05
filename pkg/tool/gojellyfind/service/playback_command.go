package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) PlaybackCommand(
	sessionIdentifier string,
	command string,
	seekPositionTicks int64,
) error {
	if sessionIdentifier == "" {
		return validation.New("session_id is required")
	}

	if command == "" {
		return validation.New("command is required")
	}

	return s.attempt(
		func() error {
			return s.client.PlaybackCommand(
				sessionIdentifier,
				command,
				seekPositionTicks,
			)
		},
	)
}
